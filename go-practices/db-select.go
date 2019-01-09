package main

import (
	"fmt"
	"database/sql"
	"github.com/go-goracle/goracle"
)

const SELECT_DICT_DESTINATION_RATES = `
-- DestinationRates
select 
	*
  from(
        select distinct sncode, spcode, strefa,apn, a ,b  
          from 
            (select tmm.sncode as sncode, 
                    tmm.spcode as spcode, 
                    (case  
                        when gvm.cgi = '*' then ''
                        else gvm.cgi
                        end ) as strefa ,
                    (case  
                        when gvm.digits = '*' then ''
                        else gvm.digits
                        end ) as apn, 
                    replace(to_char(pv.parameter_value_float * pu.ppuic),',','.') as a, 
                    pv_beat.parameter_value_float as b
               from  mpulkrim ri,
                     rate_pack_element rpe,
                     rate_pack_parameter_value pv,
                     rate_pack_parameter_value pv_beat,
                     rate_pack_parameter_value pv2,
                     rate_pack_parameter_value pv_beat2,
                     mpulkgvm gvm,
                     mpuzntab zn,
                     mputttab tt,
                     mpulktmm tmm,
                     MPULKTWM lktw,
                     mpuputab pu,
                     rateplan r,
                     mpusptab sp,
                     mpusntab sn
             where ri.ricode = tmm.ricode
               and ri.vscode = (select max(ri2.vscode) 
                                  from mpurivsd ri2 
                                 where ri2.ricode = ri.ricode)
               and gvm.gvcode  = (select unique gvcode 
                                    from mpulkrim 
                                   where ricode = ri.ricode)
               and gvm.gvcode  = (select max(gvcode) 
                                    from mpugvvsd gvsd 
                                   where gvm.gvcode = gvsd.gvcode)
               and gvm.zncode = zn.zncode
               and tmm.tmcode = r.tmcode
               and tmm.spcode = sp.spcode
               and tmm.sncode = sn.sncode
               and tmm.tmcode in (553) 
               and tmm.spcode in (546,435, 436, 440, 547, 565, 578)
               and gvm.cgi in ('1A_','1B_','2_','3_','OA_','1A2_','1B2_','ONAIR','*')
               and tmm.sncode in (111,1047)
               AND tmm.tmcode = pu.tmcode
               and tmm.vscode = (select max(rv2.vscode) 
                                   from rateplan_version rv2 
                                  where rv2.tmcode = tmm.tmcode)
               AND pu.vscode = (SELECT MAX(pu2.vscode) 
                                  FROM mpuputab pu2 
                                 WHERE pu.tmcode = pu2.tmcode)
               AND rate_type_id = 1
               AND ri.rate_pack_entry_id = rpe.rate_pack_entry_id
               AND ri.zncode = zn.zncode
               AND ri.ttcode = tt.ttcode
               AND lktw.twcode = ri.twcode
               AND lktw.vscode = (SELECT MAX(lktw2.vscode) 
                                    FROM mpulktwm lktw2 
                                   WHERE lktw2.twcode = lktw.twcode)
               AND ri.twvscode = lktw.vscode
               AND lktw.ttcode = ri.ttcode
               AND rpe.rate_pack_element_id = pv.rate_pack_element_id
               AND rpe.rate_pack_element_id = pv_beat.rate_pack_element_id
               AND rpe.rate_pack_element_id = pv2.rate_pack_element_id(+)
               AND rpe.rate_pack_element_id = pv_beat2.rate_pack_element_id(+)
               AND pv.parameter_rownum = 1
               AND pv.parameter_seqnum = /*4 */  DECODE((SELECT COUNT(*) 
                                                           FROM rate_pack_parameter_value pv_ 
                                                          WHERE pv_.rate_pack_element_id = pv.rate_pack_element_id 
                                                            AND pv_.parameter_seqnum = 4), 0, 1, 4)
               AND pv_beat.parameter_rownum = 1
               AND pv_beat.parameter_seqnum = 1
               AND pv2.parameter_rownum(+) = 2
               AND pv2.parameter_seqnum(+) = 4
               AND pv_beat2.parameter_rownum(+) = 2
               AND pv_beat2.parameter_seqnum(+) = 1
            ORDER BY tmm.tmcode,tmm.spcode, ri.ttcode, pv.rate_pack_element_id, pv.parameter_rownum,  pv.parameter_seqnum)
        )
`

func main(){
	var db *sql.DB
	var err error
	var response []string
	var result int
	var clientVersion goracle.VersionInfo

	fmt.Println("Client version:", clientVersion)

	db, err = connectToOracle()
	if err != nil{
		fmt.Println("Oracle connection error", err)
		return
	}
	response, err = ImportDictionary(db, SELECT_DICT_DESTINATION_RATES, &result)
	if err != nil{
		fmt.Println("Select execution error", err)
		return
	}
	fmt.Printf("Response: %v\n", response)
}


func connectToOracle() (*sql.DB, error){

	host := "localhost.localdomain"
	serviceName := "t2bscs.world"
	user := "orca"
	password := "Orca1234"
	port := "12521"

	/* db connection */
	connString := fmt.Sprintf("%s/%s@%s:%s/%s", user, password, host, port, serviceName)
	db, err := sql.Open("goracle", connString)
	if err !=nil{
		return nil, err
	}
	
	err = db.Ping()

	return db, err
}

func ImportDictionary(db *sql.DB, query string, result *int) ([]string, error){
	var response []string
	var rows *sql.Rows
	var row string

	rows, err := db.Query(query)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		rows.Scan(&row)
		fmt.Println(row)
		response = append(response, row)
	}
	
	return response, err
}