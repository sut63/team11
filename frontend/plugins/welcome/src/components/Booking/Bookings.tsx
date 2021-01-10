import React, { useState, useEffect } from 'react';
import ComponanceTable from '../TableClient';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content, Header, Page, pageTheme, ContentHeader,
} from '@backstage/core';
import {
  Table, TableBody, TableCell, TableRow, Typography,
  TextField, Button, withStyles, makeStyles,
  Theme, FormControl, InputLabel, MenuItem,
  FormHelperText, Select, createStyles,
  Dialog,DialogActions,DialogContent,DialogContentText,
  DialogTitle,useMediaQuery,useTheme
} from '@material-ui/core';

import { Alert } from '@material-ui/lab';

import { DefaultApi } from '../../api/apis';
import {
  EntClientEntity,
  EntServicePoint
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
  const [bookingtypes, setBookingtypes] = useState<EntBookingtype[]>(Array);
  useEffect(() => {
    const getCliententity = async () => {
      const res = await api.listCliententity({ limit: 10, offset: 0 });
      setLoading(false);
      setClients(res);
    };
    getCliententity();

    const getUsers = async () => {
      const res = await api.listUser({ limit: 10, offset: 0 });
      setLoading(false);
      setUsers(res);
    };
    getUsers();

    const getBookingtype = async () => {
      const res = await api.listServicepoint({ limit: 10, offset: 0 });
      setLoading(false);
      setBookingtypes(res);
    };
    getBookingtype();




  }, [loading]);

  const BookingDatehandleChange = (event: any) => {
    setBookingdate(event.target.value as string);
  };

  const BookingIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setBookingtypeID(event.target.value as number);
  };

  const clientIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
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
  const [bookingtypeID, setBookingtypeID] = useState(Number);
  const [clientID, setClientID] = useState(Number);
  const [userID, setUserID] = useState(Number);
 
  const handleClose = async () => {
    setOpen(false);
    const res = await api.getUser({ id: userID });
    setUsername(res);
    pname?=username?.uSERNAME;
  };
  const handleClose2 = () => {
    setOpen(false);
    setUserID(0);
    pname="";
  };
  const CreateBooking = async () => {
    const booking = {
      bookingDate: bookingdate+":00+07:00",
      timeLeft: String("2020-01-01T03:00:00+07:00"),
      bookingtype: bookingtypeID,
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
            <Button variant="contained" color="primary" onClick={handleClose2}>
              ออกจากระบบ
       </Button>
          </div>

          
        </ContentHeader>
        <ComponanceTable></ComponanceTable>
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
                  onChange={clientIDhandleChange}
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
              disabled
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
              <tr><td>เลือกประเภทของผู้ใช้งาน</td><td>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="bookingType"></InputLabel>
                <Select
                  labelId="bookingType"
                  id="bookingType"
                  value={bookingtypeID}
                  onChange={BookingIDhandleChange}
                  style={{ width: 200 }}
                >
                  {bookingtypes.map((item: EntServicePoint) => (
                    <MenuItem value={item.id}>{item.bOOKTYPENAME}</MenuItem>
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

        <div>
      
      <Dialog
        fullScreen={fullScreen}
        open={open}
        onClose={handleClose}
        aria-labelledby="responsive-dialog-title"
      >
        <DialogTitle id="responsive-dialog-title">{"กรุณาเลือก Email เพื่อเข้าสู่ระบบ"}</DialogTitle>
        <DialogContent>
          <DialogContentText>
          <table>
              <tr><td width="100">อีเมลผู้ใช้</td><td>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="email-label"></InputLabel>
                <Select
                  labelId="email-label"
                  id="email"
                  value={userID}

                  onChange={UserIDhandleChange}
                  style={{ width: 400 }}
                ><option value="">None</option>
                  {users.map((item: EntUser) => (
                    <MenuItem value={item.id} >{item.uSEREMAIL}</MenuItem>
                  ))}
                </Select>
              </FormControl>
              </td>
              </tr>
              <tr><td>ชื่อผู้ใช้</td><td>
              <FormControl
                disabled
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
                    <MenuItem value={item.id}>{item.uSERNAME}</MenuItem>
                  ))}
                </Select>
              </FormControl>
              </td>
              </tr>
            </table>
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button autoFocus onClick={handleClose2} color="primary">
            ยกเลิก
          </Button>
          <Button onClick={handleClose} color="primary" autoFocus>
            เข้าสู่ระบบ
          </Button>
        </DialogActions>
      </Dialog>
    </div>

      </Content>
    </Page>
  );
}