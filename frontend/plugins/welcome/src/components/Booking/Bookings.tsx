import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content, Header, Page, pageTheme, ContentHeader,
} from '@backstage/core';
import {
  Typography, TextField, Button, makeStyles,
  Theme, FormControl, InputLabel, MenuItem,Select, createStyles,
  Dialog,DialogActions,DialogContent,DialogContentText,
  DialogTitle,useMediaQuery,useTheme
} from '@material-ui/core';

import { Alert } from '@material-ui/lab';

import { DefaultApi } from '../../api/apis';
import {
  EntClientEntity,
  EntServicePoint,
  EntUser,
} from '../../api/models/';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '35ch',
    },
  }),
);


var pname = '';
export default function Create() {
  const [username, setUsername] = useState<EntUser>();
  const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบ VideoOnDemand' };
  const classes = useStyles();
  const api = new DefaultApi();
  const theme = useTheme();
  const fullScreen = useMediaQuery(theme.breakpoints.down('sm'));
  const [open, setOpen] = React.useState(false);
  const [loading, setLoading] = useState(true);

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [clients, setClients] = useState<EntClientEntity[]>(Array);
  const [users, setUsers] = useState<EntUser[]>(Array);
  const [servicepoint, setServicePoint] = useState<EntServicePoint[]>(Array);

  useEffect(() => {
    const getCliententity = async () => {
      const res = await api.listCliententity({ limit: 10, offset: 0 });
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

    const getServicePoint = async () => {
      const res = await api.listServicepoint({ limit: 10, offset: 0 });
      setLoading(false);
      setServicePoint(res);
    };
    getServicePoint();




  }, [loading]);

  const BookingDatehandleChange = (event: any) => {
    setBookingdate(event.target.value as string);
  };

  const ServicePointIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setServicePointID(event.target.value as number);
  };

  const ClientIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setClientID(event.target.value as number);
  };

  const UserIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUserID(event.target.value as number);
  };

  const handleClickOpen = () => {
    setOpen(true);
  };



  const [bookingdate, setBookingdate] = useState(String);
  const [timeleft, setTimeleft] = useState(String);
  const [servicepointID, setServicePointID] = useState(Number);
  const [clientID, setClientID] = useState(Number);
  const [userID, setUserID] = useState(Number);
 
  const CreateBooking = async () => {
    const booking = {
      bookingDate: bookingdate+":00+07:00",
      timeLeft: String("2020-01-01T03:00:00+07:00"),
      servicePoint: servicepointID,
      client: clientID,
      user: userID,
    };
    const cliententity = {
      sid: Number(2)
    };
    console.log(booking);
    const res2: any = await api.updateCliententity({ id: clientID, cliententity: cliententity });
    const res: any = await api.createBooking({ booking: booking });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };
  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`${profile.givenName}`}
      //subtitle="Some quick intro and links."
      ></Header>
      <Content>
        <ContentHeader title="เพิ่มข้อมูลการจอง">
          <Typography align="left" style={{ marginRight: 16, color: "#00eeff" }}>
            {pname}
          </Typography>
          <div>
          <Button variant="contained" color="primary" onClick={handleClickOpen}>
        เข้าสู่ระบบ
      </Button>
      &nbsp;&nbsp;&nbsp;
            <Button variant="contained" color="primary" >
              ออกจากระบบ
       </Button>
          </div>

          
        </ContentHeader>

        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <table>
              <tr><td width="150">เลือกเครื่องรับชม</td><td>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >

                <InputLabel id="client-label"></InputLabel>
                <Select
                  labelId="client-label"
                  id="client"
                  value={clientID}
                  onChange={ClientIDhandleChange}
                  style={{ width: 400 }}
                >
                  {clients.map((item: EntClientEntity) => (
                    <MenuItem value={item.id}>{item.cLIENTNAME}</MenuItem>
                  ))}
                </Select>
              </FormControl>
              </td>
              </tr>
              <tr><td>สมาชิกห้องสมุด</td><td>
              <FormControl
              
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="user-label"></InputLabel>
                <Select
                  labelId="user-label"
                  id="user"
                  value={userID}
                  onChange={UserIDhandleChange}
                  style={{ width: 400 }}
                >
                  {users.map((item: EntUser) => (
                    <MenuItem value={item.id}>{item.uSEREMAIL}</MenuItem>
                  ))}
                </Select>
              </FormControl>
              </td>
              </tr>
              <tr><td>เลือกจุดที่จะเข้าไปยืนยันการจอง</td><td>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="ServicePoint"></InputLabel>
                <Select
                  labelId="servicepoint"
                  id="servicepoint"
                  value={servicepointID}
                  onChange={ServicePointIDhandleChange}
                  style={{ width: 200 }}
                >
                  {servicepoint.map((item: EntServicePoint) => (
                    <MenuItem value={item.id}>{item.cOUNTERNUMBER}</MenuItem>
                  ))}
                </Select>
              </FormControl>
              </td>
              </tr>
              <tr><td>เลือกวันที่และเวลา</td><div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <TextField
                  id="deathtime"
                  label=""
                  type="datetime-local"
                  value={bookingdate}
                  onChange={BookingDatehandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </FormControl>
              </div>
              </tr>

            </table>
            <tr><td width="300">
            <Button
                onClick={() => {
                  CreateBooking();
                  
                }}
                
                variant="contained"
                color="primary"
              >
                [บันทึกการจอง]
           </Button>
              <Button
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/welcome"
                variant="contained"
              >
                กลับ
           </Button>
           </td><td>{status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  This is a success alert — check it out!
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    This is a warning alert — check it out!
                  </Alert>
                )}
            </div>
          ) : null}</td>
           </tr>
          </form>
        </div>

      </Content>
    </Page>
  );
}