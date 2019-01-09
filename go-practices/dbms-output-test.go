package main

import (
	"github.com/go-goracle/goracle"
	"database/sql"
	"fmt"
	"time"
	"context"
	"bytes"
)

const STR_ANONYMOUS_SINGLE_CONTRACT_IMPORT = `
	declare
		in_co_id number := :1;
		in_date date := :2;
		out_message varchar2(1024);
		out_result number;
		v_tmp number;
		v_sub_market number;
		v_ch_status varchar2(1);

		RETCODE_OK                          constant number := 0;
		RETCODE_INVALID_CONTRACT_ID         constant number := 101;
		RETCODE_CONTRACT_NOT_POSTPAID       constant number := 102;
		RETCODE_CONTRACT_INACTIVE           constant number := 103;
		RETCODE_RATEPLAN_UNFOUND            constant number := 104;
		RETCODE_MSISDN_UNFOUND              constant number := 105;
		RETCODE_IMSI_UNFOUND                constant number := 106;
		RETCODE_INVALID_INPUT_DATE          constant number := 107;
		RETCODE_CONTRACT_UNFOUND            constant number := 108;
		RETCODE_CONTRACT_HIST_UNFOUND       constant number := 109;
		RETCODE_PARAM_NOT_FOUND             constant number := 110;
		RETCODE_SERVICE_UNDEF_ERROR         constant number := 111;
		RETCODE_BALANCE_UNDEF_ERROR         constant number := 200;    
		SNCODE_VOICE                        constant number := 1;
		SNCODE_DATA                         constant number := 101;
		SNCODE_EUROTARIFF                   constant number := 911;    
		PARAM_EUROTARIFF                    constant number := 149;    

		type t_import_contract is record
		(
			co_id       	number(38,0),
			customer_id	  number(38,0),
			msisdn        varchar2(100 BYTE),
			imsi          varchar2(50),
			tmcode        number(38, 0),
			valid_from    date
		);
		rec_import_contract t_import_contract;
	

	begin

		--rec_import_contract := t_import_contract(null, null, null, null, null, null);
		rec_import_contract.co_id := null;
		rec_import_contract.customer_id := null;
		rec_import_contract.msisdn := null;
		rec_import_contract.imsi := null;
		rec_import_contract.tmcode := null;
		rec_import_contract.valid_from := null;


		if in_co_id is NULL then
			out_result := RETCODE_INVALID_CONTRACT_ID;
			return;
		end if;

		out_result := RETCODE_CONTRACT_UNFOUND;
		select  co_id,
				customer_id,
				subm_id
		into  rec_import_contract.co_id,
				rec_import_contract.customer_id,
				v_sub_market
		from sysadm.contract_all
		where co_id = in_co_id;

		if v_sub_market != 1 then
			out_result := RETCODE_CONTRACT_NOT_POSTPAID;
			return;
		end if;

		-- status kontraktu
		out_result := RETCODE_CONTRACT_HIST_UNFOUND;
		select ch_validfrom,
			ch_status
		into rec_import_contract.valid_from,
			v_ch_status
		from sysadm.contract_history ch1
		where ch1.co_id = in_co_id
		and ch1.ch_seqno = (select max(ch2.ch_seqno)
								from sysadm.contract_history ch2
								where ch2.co_id = ch1.co_id
								and ch2.ch_validfrom < in_date);
								
		if v_ch_status != 'a' then
			out_result := RETCODE_CONTRACT_INACTIVE;
			return;
		end if;

		-- rateplan
		out_result:= RETCODE_RATEPLAN_UNFOUND;
		select rh1.tmcode
		into rec_import_contract.tmcode
		from sysadm.rateplan_hist rh1
		where rh1.co_id = in_co_id
		and rh1.seqno = (select max(rh2.seqno)
							from sysadm.rateplan_hist rh2
						where rh2.co_id = rh1.co_id);
	
		-- msisdn
		out_result:= RETCODE_MSISDN_UNFOUND;
		select dn.dn_num
		into rec_import_contract.msisdn
		from sysadm.directory_number dn, sysadm.contr_services_cap csc
		where csc.dn_id = dn.dn_id
		and csc.co_id = in_co_id
		and csc.sncode in (SNCODE_VOICE, SNCODE_DATA)
		and csc.seqno = (select max(cap.seqno)
							from sysadm.contr_services_cap cap
							where cap.co_id = csc.co_id
							and cap.sncode = csc.sncode)
		and rownum = 1;
						
		-- imsi
		out_result:= RETCODE_IMSI_UNFOUND;
		select p.port_num
		into rec_import_contract.imsi
		from sysadm.contr_devices cd1, sysadm.port p
		where cd1.co_id = in_co_id
		and p.port_id = cd1.port_id
		and cd1.cd_seqno = (select max(cd2.cd_seqno)
								from sysadm.contr_devices cd2
								where cd2.co_id = cd1.co_id);    

		out_result := RETCODE_OK;
		out_message :=  rec_import_contract.co_id||';'||
						rec_import_contract.customer_id||';'||
						rec_import_contract.msisdn||';'||
						rec_import_contract.imsi||';'||
						rec_import_contract.tmcode||';'||
						to_char(rec_import_contract.valid_from, 'yyyymmddhh24miss');
					
		dbms_output.put_line(out_result||'|'||out_message);
					
	exception
		when others then
			null;
	end import_single_contract;`

const STR_ANONYMOUS_GET_ACTUAL_PROFILE_SERVICE = `
declare
in_co_id number := :1;
in_cursor sys_refcursor;
out_result number;
v_output varchar2(128);

type tt_import_profile_service is record
(
  co_id       	  number(38, 0),
  sncode          number(38, 0)
);

v_rec tt_import_profile_service;

RETCODE_OK                          constant number := 0;
RETCODE_INVALID_CONTRACT_ID         constant number := 101;
RETCODE_CONTRACT_NOT_POSTPAID       constant number := 102;
RETCODE_CONTRACT_INACTIVE           constant number := 103;
RETCODE_RATEPLAN_UNFOUND            constant number := 104;
RETCODE_MSISDN_UNFOUND              constant number := 105;
RETCODE_IMSI_UNFOUND                constant number := 106;
RETCODE_INVALID_INPUT_DATE          constant number := 107;
RETCODE_CONTRACT_UNFOUND            constant number := 108;
RETCODE_CONTRACT_HIST_UNFOUND       constant number := 109;
RETCODE_PARAM_NOT_FOUND             constant number := 110;
RETCODE_SERVICE_UNDEF_ERROR         constant number := 111;
RETCODE_BALANCE_UNDEF_ERROR         constant number := 200;


SNCODE_VOICE                        constant number := 1;
SNCODE_DATA                         constant number := 101;
SNCODE_EUROTARIFF                   constant number := 911;

PARAM_EUROTARIFF                    constant number := 149;    

begin
out_result := RETCODE_SERVICE_UNDEF_ERROR;

open in_cursor for 
  select ps.co_id, ps.sncode
	from sysadm.profile_service ps, sysadm.pr_serv_status_hist psh 
   where ps.co_id = in_co_id
	 and ps.co_id = psh.co_id 
	 and ps.sncode = psh.sncode 
	 and ps.status_histno = psh.histno
	 and ps.sncode in (111, 1047)
	 and psh.status = 'A'
   order by ps.sncode;
   
out_result := RETCODE_OK;

v_rec.co_id := null;
v_rec.sncode := null;

v_output := out_result;

loop
  fetch in_cursor into
	v_rec.co_id,
	v_rec.sncode;
   exit when in_cursor%notfound;
	v_output := v_output||'|'||v_rec.co_id||';'||v_rec.sncode;
end loop;
close in_cursor;
dbms_output.put_line(v_output);   


exception
when others then
  null;       
end;`

func main(){
	var testDb *sql.DB
	var err error
	var clientVersion goracle.VersionInfo
	var serverVersion goracle.VersionInfo

	connString := fmt.Sprintf("%s/%s@%s:%s/%s", "orca", "Orca1234", "localhost.localdomain", "1521", "XE")

	if testDb, err = sql.Open("goracle", connString); err != nil {
		fmt.Printf("ERROR: %+v\n", err)
		return
		//panic(err)
	}	

	if testDb != nil {
		if clientVersion, err = goracle.ClientVersion(testDb); err != nil {
			fmt.Printf("ERROR: %+v\n", err)
			return
		}
		if serverVersion, err = goracle.ServerVersion(testDb); err != nil {
			fmt.Printf("ERROR: %+v\n", err)
			return
		}
		fmt.Println("Server:", serverVersion)
		fmt.Println("Client:", clientVersion)
	}	

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := testDb.Conn(ctx)

	if err != nil {
		fmt.Printf("ERROR: %+v\n", err)
	}
	defer conn.Close()
	if err := goracle.EnableDbmsOutput(ctx, conn); err != nil {
		fmt.Printf("ERROR: %+v\n", err)
	}

	if _, err := conn.ExecContext(ctx, STR_ANONYMOUS_SINGLE_CONTRACT_IMPORT, 8110388, time.Now()); err != nil {
		fmt.Printf("ERROR: %+v\n", err)
	}

	var buf bytes.Buffer
	if err := goracle.ReadDbmsOutput(ctx, &buf, conn); err != nil {
		fmt.Printf("ERROR: %+v\n", err)	
	}	
	fmt.Printf("read: [%s]\n", buf.String())

	if _, err := conn.ExecContext(ctx, STR_ANONYMOUS_GET_ACTUAL_PROFILE_SERVICE, 8110388); err != nil {
		fmt.Printf("ERROR: %+v\n", err)
	}
	if err := goracle.ReadDbmsOutput(ctx, &buf, conn); err != nil {
		fmt.Printf("ERROR: %+v\n", err)	
	}
	fmt.Printf("read: [%s]\n", buf.String())

}

