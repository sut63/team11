import React, { useState, useEffect } from 'react';
import SaveAltIcon from '@material-ui/icons/SaveAlt';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
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
  FormControl, InputLabel, MenuItem,Select,
  Typography,
  Grid,
  makeStyles,
  Button,
  FormHelperText,
} from '@material-ui/core';

import { Alert, AlertTitle } from '@material-ui/lab';

import { DefaultApi } from '../../api/apis';
import {
  EntClientEntity,
  EntServicePoint,
  EntUser,
} from '../../api/models/';

const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
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
    flexGrow: 1,
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
}));


export default function Create() {
  const name = JSON.parse(String(localStorage.getItem("userName")));
  const userName ="ยินดีต้อนรับ "+name
  const classes = useStyles();
  const api = new DefaultApi();

  const [loading, setLoading] = useState(true);

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [clients, setClients] = useState<EntClientEntity[]>(Array);
  const [users, setUsers] = useState<EntUser[]>(Array);
  const [servicepoint, setServicePoint] = useState<EntServicePoint[]>(Array);

  const idString = JSON.parse(String(localStorage.getItem("userID")));
  const idInt = parseInt(idString);

  useEffect(() => {
    const getCliententity = async () => {
      const res = await api.listCliententity();
      setLoading(false);
      setClients(res);
    };
    getCliententity();

    const getUsers = async () => {
      const res = await api.listUser();
      setLoading(false);
      setUsers(res);
    };
    getUsers();
    
    
    setUserID(idInt);

    const getServicePoint = async () => {
      const res = await api.listServicepoint({ limit: 10, offset: 0 });
      setLoading(false);
      setServicePoint(res);
    };
    getServicePoint();




  }, [loading]);

  const ServicePointIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setServicePointID(event.target.value as number);
  };

  const ClientIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setClientID(event.target.value as number);
  };

  const [servicepointID, setServicePointID] = useState(Number);
  const [clientID, setClientID] = useState(Number);
  const [userID, setUserID] = useState(Number);
 
  const CreateBooking = async () => {

    if ((servicepointID != 0) && (clientID != 0)) {
      const resC = await api.listCliententity();
      setClients(resC);
      const booking = {
        servicePoint: servicepointID,
        client: clientID,
        user: userID,
      };
      const cliententity = {
        sid: Number(2)
      };

      const res: any = await api.createBooking({ booking: booking });
      setStatus(true);
      if (res.id != '') {
        await api.updateCliententity({ id: clientID, cliententity: cliententity });
        setAlert(true);
      }
    } else {
      setStatus(true);
      setAlert(false);
    }
    const timer = setTimeout(() => {
      setStatus(false);
      window.location.reload(false);
    }, 10000);
  };
  const resetLocalStorage = async () => {
    localStorage.setItem("userID", JSON.stringify(null))
    localStorage.setItem("role", JSON.stringify(null))
    localStorage.setItem("valid", JSON.stringify(null))
    localStorage.setItem("userName", JSON.stringify(null))
    window.location.href = "/"
  }

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

            <InfoCard title="กรุณากรอกข้อมูลเพื่อจองเครื่องรับชม VideoOnDemand 💻">
              <Typography variant="body1" gutterBottom>
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;🖱 Log off ออกจากเครื่องรับชมเมื่อเลิกใช้งาน และส่งคืนหูฟังที่เคาน์เตอร์
                (After using the video on demand machines, please log off. Then kindly return the headphone to the staff at the Information Counter.)<br/><br/>
   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;👨‍💻 ผู้ใช้บริการต้องรับผิดชอบต่ออุปกรณ์ที่ชำรุดเสียหาย (Library users are liable to any damage incurred or lost.)
              </Typography>
              <br/><br/>
              <Grid container alignItems="center" direction="column">
              <FormControl required className={classes.formControl}>
              <InputLabel id="demo-simple-select-required-label">Client</InputLabel>
        <Select
          labelId="demo-simple-select-required-label"
          id="demo-simple-select-required"
          value={clientID}
          className={classes.selectEmpty}
          onChange={ClientIDhandleChange}
          style={{ width: 400 }}
        >
          {clients.filter((filter:any) => filter.edges.state.sTATUSNAME == "Available").map((item: EntClientEntity) => (
                    <MenuItem value={item.id}>{item.cLIENTNAME}</MenuItem>
                  ))}
        </Select>
        <FormHelperText>Required</FormHelperText>
      </FormControl>
      <br/>
      <FormControl required className={classes.formControl}>
              <InputLabel id="demo-simple-select-required-label">Library Member</InputLabel>
        <Select
          disabled = {true}
          labelId="demo-simple-select-required-label"
          id="demo-simple-select-required"
          value={userID}
          className={classes.selectEmpty}
          style={{ width: 400 }}
        >
          {users.map((item: EntUser) => (
                    <MenuItem value={item.id}>{item.uSEREMAIL}</MenuItem>
                  ))}
        </Select>

      </FormControl>
      <br/>
      <FormControl required className={classes.formControl}>
              <InputLabel id="demo-simple-select-required-label">Counter</InputLabel>
        <Select
          labelId="demo-simple-select-required-label"
          id="demo-simple-select-required"
          value={servicepointID}
          className={classes.selectEmpty}
          onChange={ServicePointIDhandleChange}
          style={{ width: 400 }}
        >
          {servicepoint.map((item: EntServicePoint) => (
                    <MenuItem value={item.id}>{item.cOUNTERNUMBER}</MenuItem>
                  ))}
        </Select>
        <FormHelperText>Required</FormHelperText>
      </FormControl>
                    <Button
                    variant="contained"
                    color="secondary"
                    className={classes.button}
                    startIcon={<SaveAltIcon />}
                    onClick={() => {
                      CreateBooking();
                      setClientID(0);
                      setServicePointID(0);
                    }}
                  >
                    บันทึกการจอง
                  </Button>
                  <Grid container justify="center" item xs={12}>
        {status ? (
           <div>
             {alert ? (
               <Alert severity="success" style={{ width: 400 }} onClose={() => { setStatus(false); window.location.reload(false);}}>
               <AlertTitle>Success</AlertTitle>
               <div>
               ทำการบันทึกข้อมูลการจองสำเร็จ — <strong>🎉</strong>
               </div>
               <br/>
             </Alert> 
             ) : (
              <Alert severity="error" style={{ width: 400 }} onClose={() => { setStatus(false); window.location.reload(false);}}>
               <AlertTitle>Error</AlertTitle >
               <div>
               บันทึกข้อมูลผิดพลาด — <strong>❌</strong>
               </div><br/>
             </Alert>
             )}
           </div>
         ) : null}
         </Grid>
      </Grid>
            </InfoCard>
          </Grid>
          <Grid item xs={12} md={6}>
          <InfoCard title="ตารางเครื่องที่ว่าง">

            </InfoCard>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
}