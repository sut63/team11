import React, { useState } from 'react';
import SaveAltIcon from '@material-ui/icons/Search';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import { DataGrid, ColDef } from '@material-ui/data-grid';
import ClearAllIcon from '@material-ui/icons/ClearAll';
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
  FormControl, InputLabel, MenuItem, Select,
  Typography,
  Grid,
  TextField,
  makeStyles,
  Button,
} from '@material-ui/core';

import { Alert, AlertTitle } from '@material-ui/lab';

import { DefaultApi } from '../../api/apis';

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
  const [load, setLoad] = useState(false);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);


  const [filter, setFilter] = useState("Or");
  const handleSelect = (event: React.ChangeEvent<{ value: unknown }>) => {
    setFilter(event.target.value as string);
  };

  const handleChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as keyof typeof Search;
    const { value } = event.target;
    setBooking({ ...bookingField, [id]: value });
  };

  const [rows, setRows] = useState([]);
  var r: any = []

  const SearchBooking = async () => {

    const BI = Number(bookingField.searchID)
    const BU = bookingField.searchUserName
    const BC = bookingField.searchClient
    const BP = bookingField.searchPhoneNumber
    const F = filter
    console.log(F)

    if (BI == 0 && BU == "" && BC == "" && BP == "") {
      setRows([])
      setLoad(false)
    } else {
      const res = await api.searchBooking({   //ทำการส่งข้อมูลไปค้นหาที่ backend
        userid: BI,
        username: BU,
        clientname: BC,
        phonenumber: BP,
        operator: F,
      });
      console.log(res)

      //ทำการนำข้อมูลใส่ตาราง
      res.map((row: any) => (
        r.push({
          id: row.id,
          userName: row.edges.usedby.uSERNAME,
          client: row.edges.using.cLIENTNAME,
          phoneNumber: row.pHONENUMBER,
          userNumber: row.uSERNUMBER,
          borrowItem: row.bORROWITEM,
          bookingDate: row.bOOKINGDATE,
          timeOut: row.tIMELEFT,
          servicePoint: row.edges.getservice.cOUNTERNUMBER
        })
      ))
      setRows(r)

    }

    if (r.length == 0) {
      setStatus(true);
      setAlert(false);
    }
    else {
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
    { field: 'id', type: 'number', headerName: 'ID', width: 90 },
    { field: 'userName', headerName: 'ชื่อผู้จอง ', width: 200 },
    { field: 'client', headerName: 'ชื่อเครื่องรับชม', width: 180 },
    { field: 'phoneNumber', headerName: 'หมายเลขโทรศัพท์ผู้จอง', width: 200 },
    { field: 'userNumber', type: 'number', headerName: 'จำนวนผู้ใช้งาน', width: 145 },
    { field: 'borrowItem', type: 'number', headerName: 'จำนวนอุปกรณ์ที่ยืม', width: 170 },
    { field: 'bookingDate', headerName: 'วันเวลาที่จอง', width: 210 },
    { field: 'timeOut', headerName: 'วันที่หมดเวลา', width: 210 },
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
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;🖱 การค้นหา ผู้ใช้สามารถค้นหาข้อมููลได้ตามช่องที่ใส่ข้อมูลด้านล่างนี้ โดยเมื่อใส่ข้อมูลไปแล้ว ผุ้ใช้สามารถนำคำค้นหาทั้งหมด มาเชื่อมกันได้ด้วย LinkFilter ทั้งแบบ "And" และ "Or"
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
                        <MenuItem value={"Or"}>Or</MenuItem>
                        <MenuItem value={"And"}>And</MenuItem>
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
                        <MenuItem value={"Or"}>Or</MenuItem>
                        <MenuItem value={"And"}>And</MenuItem>
                      </Select>
                    </FormControl>
                  </Grid>
                </Grid>

                {/* ค้นหา Client */}
                <Grid container alignItems="center" justify="center" item xs={12}>
                  <Grid item xs={3}>
                    <Typography variant="button" display="block" gutterBottom>
                      ค้นหาด้วย หมายเลขเครื่องรับชม
                  </Typography>
                  </Grid>
                  <Grid>
                    <TextField id="searchClient" onChange={handleChange} type="number" InputProps={{ inputProps: { min: 1 } }} value={bookingField.searchClient} label="ใส่ หมายเลขเครื่อง" style={{ width: 225 }} variant="outlined" />
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
                        <MenuItem value={"Or"}>Or</MenuItem>
                        <MenuItem value={"And"}>And</MenuItem>
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
                    <TextField id="searchPhoneNumber" value={bookingField.searchPhoneNumber} onChange={handleChange} label="ใส่หมายเลขโทรศัพท์" style={{ width: 225 }} variant="outlined" />
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
                    }}
                  >
                    เคลียร์ข้อมูล
                  </Button>
                </Grid>

                <Grid container justify="center" item xs={12}>


                  {status ? (
                    <div>
                      {alert ? (
                        <Alert severity="success" style={{ width: 400 }} onClose={() => { setStatus(false) }}>
                          <AlertTitle>Success</AlertTitle>
                          <div>
                            พบข้อมูลที่ค้นหา — <strong>🎉</strong>
                          </div>
                          <br />
                        </Alert>
                      ) : (
                          <Alert severity="error" style={{ width: 400 }} onClose={() => { setStatus(false) }}>
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