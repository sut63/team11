import React, { useState, useEffect } from 'react';
import SaveAltIcon from '@material-ui/icons/Search';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import { DataGrid, ColDef } from '@material-ui/data-grid';
import ClearAllIcon from '@material-ui/icons/ClearAll';
import ComponanceTable from './Table/TablesBookings';
import {
  Content,
  InfoCard,
  Header,
  HomepageTimer,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import {
  Paper,
  FormControl, InputLabel, MenuItem, Select,
  Typography,
  Grid,
  TextField,
  makeStyles,
  Button,
  FormHelperText,
} from '@material-ui/core';

import { Alert, AlertTitle } from '@material-ui/lab';

import { DefaultApi } from '../../api/apis';
import {
  EntBooking,
} from '../../api/models';

const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  paperG: {
    padding: theme.spacing(2),
    textAlign: 'center',
    color: theme.palette.text.secondary,
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%',
    marginTop: theme.spacing(0),
  },
  submit: {
    margin: theme.spacing(1, 0, 2),
  },
  formControl: {
    margin: theme.spacing(1),
    minWidth: 120,
  },
  button: {
    margin: theme.spacing(1),
  },
  img: {
    margin: 'auto',
    display: 'block',
    maxWidth: '100%',
    maxHeight: '100%',
  },
  root: {
    ...theme.typography.button,
    backgroundColor: theme.palette.background.paper,
    padding: theme.spacing(1),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
}));

interface operatorField {
  operatorsID: string;
  operatorsName: string;
  operatorsClient: string;
  operatorsPhone: string;
}
interface bookingField {
  searchID: number;
  searchUserName: string;
  searchClient: string;
  searchPhoneNumber: string;
}
export default function Search() {
  const name = JSON.parse(String(localStorage.getItem("userName")));
  const userName = "ยินดีต้อนรับ " + name
  const classes = useStyles();
  const api = new DefaultApi();

  //Interface
  const [bookingField, setBooking] = useState<Partial<bookingField>>({
    searchID: 0,
    searchUserName: "",
    searchClient: "",
    searchPhoneNumber: ""
  });
  const [operatorField, setOperator] = useState<Partial<operatorField>>({
    operatorsID: "=",
    operatorsName: "contains",
    operatorsClient: "contains",
    operatorsPhone: "contains"
  });

  const [loading, setLoading] = useState(true);
  const [load, setLoad] = useState(false);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [bookings, setBookings] = useState<EntBooking[]>(Array);



  useEffect(() => {
    const getBookings = async () => {
      const res = await api.listBooking();
      setLoading(false);
      setBookings(res);
    };
    getBookings();



  }, [loading]);


  const [phoneNumberError, setPhoneNumberError] = useState('');

  const [alertMessage, setAlertMessage] = useState('');

  const [filter, setFilter] = useState("||");
  const handleSelect = (event: React.ChangeEvent<{ value: unknown }>) => {
    setFilter(event.target.value as string);
  };

  const handleChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as keyof typeof Search;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setBooking({ ...bookingField, [id]: value });
  };
  const handleOperator = (event: React.ChangeEvent<{ name?: string; value: any }>) => {
    const name = event.target.name as keyof typeof operatorField;

    const { value } = event.target;
    setOperator({ ...operatorField, [name]: value });
  };

  const validatePhoneNumber = (val: string) => {
    return val.match("[0][689]\\d{8}") && val.length == 10;
  }

  const checkPattern = (id: string, value: any) => {
    switch (id) {
      case 'searchPhoneNumber':
        validatePhoneNumber(value) ? setPhoneNumberError('') : setPhoneNumberError('หมายเลขโทรศัพท์จะต้องขึ้นต้นด้วย 06,08,09 และตามด้วยเลข 0-9 อีก 8 ตัว');
        return;
      default:
        return;
    }
  }
  const [rows, setRows] = useState([]);
  const r: any = []

  const SearchBooking = async () => {

    setRows([])
    const BI = Number(bookingField.searchID)
    const BU = bookingField.searchUserName
    const BC= bookingField.searchClient
    const BP = bookingField.searchPhoneNumber
    const OI = operatorField.operatorsID
    const OU = operatorField.operatorsName
    const OC= operatorField.operatorsClient
    const OP = operatorField.operatorsPhone
    const F = filter

    function compareSym(sym:any, value:number,search:number) {
      if(search == 0){
        return null
      }
      else if(sym=="="){
        return value==search
      }
      else if(sym=="!="){
        return value != search
      }
      else if(sym==">"){
        return value > search
      }
      else if(sym==">="){
        return value >= search
      }
      else if(sym=="<"){
        return value < search
      }
      else if(sym=="<="){
        return value <= search
      }
    }
    function compareString(sym:any, value:string,search:any) {
      if(search==""){
        return null
      }
      else if(sym=="contains"){
        return value.indexOf(search) !== -1
      }
      else if(sym=="equals"){
        return value == search
      }
      else if(sym=="start with"){
        return value.startsWith(search)
      }
      else if(sym=="end with"){
        return value.endsWith(search)
      }
    }
    function compareSymAnd(sym:any, value:number,search:number) {
      if(search == 0){
        return true
      }
      else if(sym=="="){
        return value==search
      }
      else if(sym=="!="){
        return value != search
      }
      else if(sym==">"){
        return value > search
      }
      else if(sym==">="){
        return value >= search
      }
      else if(sym=="<"){
        return value < search
      }
      else if(sym=="<="){
        return value <= search
      }
    }
    function compareStringAnd(sym:any, value:string,search:any) {
      if(search==""){
        return true
      }
      else if(sym=="contains"){
        return value.indexOf(search) !== -1
      }
      else if(sym=="equals"){
        return value == search
      }
      else if(sym=="start with"){
        return value.startsWith(search)
      }
      else if(sym=="end with"){
        return value.endsWith(search)
      }
    }
    if (BI==0 && BU==""&&BC==""&&BP=="") {
      setRows([])
      setLoad(false)
    }else{
      if(F=="||"){
        bookings.filter((filter: any) => 
        compareSym(OI,filter.id,BI) || 
        compareString(OU,filter.edges.usedby.uSERNAME,BU) ||
        compareString(OC,filter.edges.using.cLIENTNAME,BC) ||
        compareString(OP,filter.pHONENUMBER,BP)).map((row: any) => (
          r.push({
            id: row.id,
            userName: row.edges.usedby.uSERNAME,
            client: row.edges.using.cLIENTNAME,
            phoneNumber: row.pHONENUMBER,
            userNumber: row.uSERNUMBER,
            borrowItem: row.bORROWITEM,
            servicePoint: row.edges.getservice.cOUNTERNUMBER
          })
        ))
        setRows(r)
      }
      else{
        bookings.filter((filter: any) => 
        compareSymAnd(OI,filter.id,BI) &&
        compareStringAnd(OU,filter.edges.usedby.uSERNAME,BU) &&
        compareStringAnd(OC,filter.edges.using.cLIENTNAME,BC) &&
        compareStringAnd(OP,filter.pHONENUMBER,BP)).map((row: any) => (
          r.push({
            id: row.id,
            userName: row.edges.usedby.uSERNAME,
            client: row.edges.using.cLIENTNAME,
            phoneNumber: row.pHONENUMBER,
            userNumber: row.uSERNUMBER,
            borrowItem: row.bORROWITEM,
            servicePoint: row.edges.getservice.cOUNTERNUMBER
          })
        ))
        setRows(r)
      }

    }
    console.log(rows)
    if(rows.length==0){
      setStatus(true);
      setAlert(false);
    }
    else{
      setStatus(true);
      setAlert(true);
    }


    const timer = setTimeout(() => {
      // window.location.reload(false);
    }, 10000);
  };
  const resetLocalStorage = async () => {
    localStorage.setItem("userID", JSON.stringify(null))
    localStorage.setItem("role", JSON.stringify(null))
    localStorage.setItem("valid", JSON.stringify(null))
    localStorage.setItem("userName", JSON.stringify(null))
    window.location.href = "/"
  }

  const columns: ColDef[] = [
    { field: 'id', type: 'number', headerName: 'ID', width: 70 },
    { field: 'userName', headerName: 'ชื่อผู้จอง ', width: 200 },
    { field: 'client', headerName: 'ชื่อเครื่องรับชม', width: 180 },
    { field: 'phoneNumber', headerName: 'หมายเลขโทรศัพท์ผู้จอง', width: 200 },
    { field: 'userNumber', type: 'number', headerName: 'จำนวนผู้ใช้งาน', width: 145 },
    { field: 'borrowItem', type: 'number', headerName: 'จำนวนอุปกรณ์ที่ยืม', width: 170 },
    { field: 'servicePoint', headerName: 'ติดต่อรับเครื่องได้ที่', width: 170 },
  ];

  return (

    <Page theme={pageTheme.home}>
      <Header
        title={'Welcome to VideoOnDemand System'}
        subtitle="The Center for Library Resources and Educational Media"
      >
        <HomepageTimer />
      </Header>
      <Content>
        <ContentHeader title={userName}>
          <Button
            // disabled={LogoutBtn}
            variant="contained"
            color="primary"
            className={classes.button}
            startIcon={<LockOutlinedIcon />}
            onClick={() => {
              resetLocalStorage();
            }}>
            ล็อกเอ้าท์
          </Button>
        </ContentHeader>

        <Grid container justify="center" >

          <Grid item xs={12} md={6}>

            <InfoCard title="กรอกข้อมูลสำหรับค้นหาการจองเครื่องรับชม VideoOnDemand 🔎">
              <Typography variant="body1" gutterBottom>
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;🖱 Log off ออกจากเครื่องรับชมเมื่อเลิกใช้งาน และส่งคืนหูฟังที่เคาน์เตอร์
                (After using the video on demand machines, please log off. Then kindly return the headphone to the staff at the Information Counter.)<br />
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;👨‍💻 ผู้ใช้บริการต้องรับผิดชอบต่ออุปกรณ์ที่ชำรุดเสียหาย (Library users are liable to any damage incurred or lost.)
              </Typography>
              <br />
              <Grid container alignItems="center" justify="center" spacing={3} >
                {/* ค้นหา ID */}
                <Grid container alignItems="center" justify="center" item xs={12}>
                  <Grid item xs={3}>
                    <Typography variant="button" display="block" gutterBottom>
                      ค้นหาด้วย ID
                  </Typography>
                  </Grid>
                  <Grid>
                    <TextField id="searchID" onChange={handleChange} value={bookingField.searchID} label="ใส่ ID ที่ต้องหารค้นหา" type="number" style={{ width: 225 }} InputProps={{ inputProps: { min: 1 } }} variant="outlined" />
                  </Grid>
                  <Grid>

                    <FormControl variant="outlined" className={classes.formControl}>
                      <InputLabel id="operatorsIDLabel">Operators</InputLabel>
                      <Select
                        labelId="operatorsIDLabel"
                        name="operatorsID"
                        onChange={handleOperator}
                        value={operatorField.operatorsID}
                        label="Operators"
                      >
                        <MenuItem value={"="}>=</MenuItem>
                        <MenuItem value={"!="}>!=</MenuItem>
                        <MenuItem value={">"}>{'>'}</MenuItem>
                        <MenuItem value={">="}>{'>'}=</MenuItem>
                        <MenuItem value={"<"}>{'<'}</MenuItem>
                        <MenuItem value={"<="}>{'<'}=</MenuItem>

                      </Select>
                    </FormControl>
                  </Grid>
                  <Grid item xs={2}>
                    <FormControl variant="outlined" className={classes.formControl}>
                      <InputLabel id="LinkFilter">LinkFilter</InputLabel>
                      <Select
                        labelId="LinkFilter"
                        id="LinkFilter"
                        value={filter}
                        onChange={handleSelect}
                        label="LinkFilter"
                      >
                        <MenuItem value={"||"}>Or</MenuItem>
                        <MenuItem value={"&&"}>And</MenuItem>
                      </Select>
                    </FormControl>
                  </Grid>
                </Grid>

                {/* ค้นนหา User */}
                <Grid container alignItems="center" justify="center" item xs={12}>
                  <Grid item xs={3} >
                    <Typography variant="button" display="block" gutterBottom>
                      ค้นหาด้วย ชื่อ-สกุล ผู้จอง
                  </Typography>
                  </Grid>
                  <Grid>
                    <TextField id="searchUserName" onChange={handleChange} value={bookingField.searchUserName} label="ใส่ ชื่อ-สกุล ที่ต้องหารค้นหา" style={{ width: 225 }} variant="outlined" />
                  </Grid>
                  <Grid>
                    <FormControl variant="outlined" className={classes.formControl}>
                      <InputLabel id="operatorsNameLabel">Operators</InputLabel>
                      <Select
                        labelId="operatorsNameLabel"
                        name="operatorsName"
                        onChange={handleOperator}
                        value={operatorField.operatorsName}
                        label="Operators"
                      >
                        <MenuItem value={"contains"}>contains</MenuItem>
                        <MenuItem value={"equals"}>equals</MenuItem>
                        <MenuItem value={"start with"}>start with</MenuItem>
                        <MenuItem value={"end with"}>end with</MenuItem>


                      </Select>
                    </FormControl>
                  </Grid>
                  <Grid item xs={2}>
                    <FormControl variant="outlined" className={classes.formControl}>
                      <InputLabel id="LinkFilterLabel">LinkFilter</InputLabel>
                      <Select
                        labelId="LinkFilterLabel"
                        id="LinkFilter"
                        disabled={true}
                        value={filter}
                        label="LinkFilter"
                      >
                        <MenuItem value={"||"}>Or</MenuItem>
                        <MenuItem value={"&&"}>And</MenuItem>
                      </Select>
                    </FormControl>
                  </Grid>
                </Grid>

                {/* ค้นหา Client */}
                <Grid container alignItems="center" justify="center" item xs={12}>
                  <Grid item xs={3}>
                    <Typography variant="button" display="block" gutterBottom>
                      ค้นหาด้วย ชื่อเครื่องรับชม
                  </Typography>
                  </Grid>
                  <Grid>
                    <TextField id="searchClient" onChange={handleChange} value={bookingField.searchClient} label="ใส่ ชื่อเครื่อง ที่ต้องหารค้นหา" style={{ width: 225 }} variant="outlined" />
                  </Grid>
                  <Grid>
                    <FormControl variant="outlined" className={classes.formControl}>
                      <InputLabel id="operatorsClientLabel">Operators</InputLabel>
                      <Select
                        labelId="operatorsClientLabel"
                        name="operatorsClient"
                        onChange={handleOperator}
                        value={operatorField.operatorsClient}

                        label="Operators"
                      >
                        <MenuItem value={"contains"}>contains</MenuItem>
                        <MenuItem value={"equals"}>equals</MenuItem>
                        <MenuItem value={"start with"}>start with</MenuItem>
                        <MenuItem value={"end with"}>end with</MenuItem>

                      </Select>
                    </FormControl>
                  </Grid>

                  <Grid item xs={2}>
                    <FormControl variant="outlined" className={classes.formControl}>
                      <InputLabel id="LinkFilterLabel">LinkFilter</InputLabel>
                      <Select
                        labelId="LinkFilterLabel"
                        id="LinkFilter"
                        disabled={true}
                        value={filter}
                        label="LinkFilter"
                      >
                        <MenuItem value={"||"}>Or</MenuItem>
                        <MenuItem value={"&&"}>And</MenuItem>
                      </Select>
                    </FormControl>
                  </Grid>
                </Grid>

                {/* ค้นหา PhoneNumber */}
                <Grid container alignItems="center" justify="center" item xs={12}>
                  <Grid item xs={3}>
                    <Typography variant="button" display="block" gutterBottom>
                      ค้นหาด้วย <br />หมายเลขโทรศัพท์ผู้จอง
                  </Typography>
                  </Grid>
                  <Grid>
                    <TextField  id="searchPhoneNumber" value={bookingField.searchPhoneNumber} onChange={handleChange} label="ใส่หมายเลขโทรศัพท์" style={{ width: 225 }} variant="outlined" />
                  </Grid>
                  <Grid>
                    <FormControl variant="outlined" className={classes.formControl}>
                      <InputLabel id="operatorsPhoneLabel">Operators</InputLabel>
                      <Select
                        labelId="operatorsPhoneLabel"
                        name="operatorsPhone"
                        onChange={handleOperator}
                        value={operatorField.operatorsPhone}

                        label="Operators"
                      >
                        <MenuItem value={"contains"}>contains</MenuItem>
                        <MenuItem value={"equals"}>equals</MenuItem>
                        <MenuItem value={"start with"}>start with</MenuItem>
                        <MenuItem value={"end with"}>end with</MenuItem>

                      </Select>
                    </FormControl>
                  </Grid>

                  <Grid item xs={2}>

                  </Grid>
                </Grid>
                <Grid container justify="center" item xs={12}>
                  <Button
                    variant="contained"
                    color="primary"
                    className={classes.button}
                    startIcon={<SaveAltIcon />}
                    onClick={() => {
                      SearchBooking();
                    }}
                  >
                    ค้นหา
                  </Button>
                  <Button
                    variant="contained"
                    color="secondary"
                    className={classes.button}
                    startIcon={<ClearAllIcon />}
                    onClick={() => {
                      setBooking({
                        searchID: 0,
                        searchUserName: "",
                        searchClient: "",
                        searchPhoneNumber: ""
                      });
                      setOperator({
                        operatorsID: "=",
                        operatorsName: "contains",
                        operatorsClient: "contains",
                        operatorsPhone: "contains"
                      })
                    }}
                  >
                    เคลียร์ข้อมูล
                  </Button>
                </Grid>

                <Grid container justify="center" item xs={12}>


                  {status ? (
                    <div>
                      {alert ? (
                        <Alert severity="success" style={{ width: 400 }} onClose={() => { setStatus(false)}}>
                          <AlertTitle>Success</AlertTitle>
                          <div>
                            พบข้อมูลที่ค้นหา — <strong>🎉</strong>
                          </div>
                          <br />
                        </Alert>
                      ) : (
                          <Alert severity="error" style={{ width: 400 }} onClose={() => { setStatus(false)}}>
                            <AlertTitle>Error</AlertTitle >
                            <div>
                              ไม่พบข้อมูลที่ค้นหา — <strong>❌</strong>
                            </div><br />
                          </Alert>
                        )}
                    </div>
                  ) : null}
                </Grid>
              </Grid>
            </InfoCard>
          </Grid>
          <Grid item xs={12} md={6}>
            <InfoCard title="ผลการค้นหาการจอง">
              <div style={{ height: 510, width: '100%' }}>
                <DataGrid rows={rows}
                  hideFooter={true} loading={load} disableColumnMenu={true} columns={columns} />
              </div>
            </InfoCard>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
}