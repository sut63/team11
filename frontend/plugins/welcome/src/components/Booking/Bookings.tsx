import React, { useState, useEffect } from 'react';
import SaveAltIcon from '@material-ui/icons/SaveAlt';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import ComponanceTable from './Table/TablesClient';
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
    '& > *': {
      margin: theme.spacing(1),
      width: '25ch',
    },
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
  /**
     * 
     * @type {string}
     * @memberof bookingField
     */
  phone_number?: string;
  /**
     * 
     * @type {number}
     * @memberof bookingField
     */
  user_number?: number;
  /**
     * 
     * @type {string}
     * @memberof bookingField
     */
  borrow_item?: string;
}
export default function Create() {
  const name = JSON.parse(String(localStorage.getItem("userName")));
  const userName = "ยินดีต้อนรับ " + name
  const classes = useStyles();
  const api = new DefaultApi();

  const [loading, setLoading] = useState(true);
  const [bookingField, setBooking] = useState<Partial<bookingField>>({});
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

  const [phoneNumberError, setPhoneNumberError] = useState('');
  const [userNumberError, setUserNumberError] = useState('');
  const [borrowItemError, setBorrowItemError] = useState('');
  const [alertMessage, setAlertMessage] = useState('');

  const handleChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as keyof typeof Create;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setBooking({ ...bookingField, [id]: value });
  };

  const validatePhoneNumber = (val: string) => {
    return val.match("[0][689]\\d{8}") && val.length == 10;
  }
  const validateUserNumber = (val: number) => {
    return val > 6 ? false : true
  }
  const validateBorrowItem = (val: number) => {
    return val > 6 ? false : true
  }
  const checkPattern = (id: string, value: any) => {
    switch (id) {
      case 'phone_number':
        validatePhoneNumber(value) ? setPhoneNumberError('') : setPhoneNumberError('หมายเลขโทรศัพท์จะต้องขึ้นต้นด้วย 06,08,09 และตามด้วยเลข 0-9 อีก 8 ตัว');
        return;
      case 'user_number':
        validateUserNumber(value) ? setUserNumberError('') : setUserNumberError('จำนวนผู้ใช้ต้อง ตั้งแต่ 1-6 คน');
        return;
      case 'borrow_item':
        validateBorrowItem(value) ? setBorrowItemError('') : setBorrowItemError('จำนวน HeadSet ต้อง ตั้งแต่ 1-6 ชิ้น');
        return;
      default:
        return;
    }
  }

  const checkCaseSaveError = (field: string) => {
    switch (field) {
      case 'PHONE_NUMBER':
        setAlertMessage("error ข้อมูล field phone_number ผิด");
        return;
      case 'USER_NUMBER':
        setAlertMessage("error ข้อมูล field user_number ผิด");
        return;
      case 'BORROW_ITEM':
        setAlertMessage("error ข้อมูล field borrow_item ผิด");
        return;
      default:
        setAlertMessage("บันทึกไม่สำเร็จ");
        return;
    }
  }

  const [servicepointID, setServicePointID] = useState(Number);
  const [clientID, setClientID] = useState(Number);
  const [userID, setUserID] = useState(Number);

  const CreateBooking = async () => {
    console.log(bookingField)
    const resC = await api.listCliententity();
    setClients(resC);
    const booking = {
      servicePoint: servicepointID,
      client: clientID,
      user: userID,
      phoneNumber: String(bookingField.phone_number),
      userNumber: Number(bookingField.user_number),
      borrowItem: Number(bookingField.borrow_item),
    };
    const cliententity = {
      sid: Number(2)
    };
    console.log(booking)
    const apiUrl = 'http://localhost:8080/api/v1/bookings';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(booking),
    };

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          setStatus(true);
          api.updateCliententity({ id: clientID, cliententity: cliententity });
          setAlert(true);
        } else {
          setStatus(true);
          setAlert(false);
          checkCaseSaveError(data.error.Name)
        }
      });
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
                (After using the video on demand machines, please log off. Then kindly return the headphone to the staff at the Information Counter.)<br />
   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;👨‍💻 ผู้ใช้บริการต้องรับผิดชอบต่ออุปกรณ์ที่ชำรุดเสียหาย (Library users are liable to any damage incurred or lost.)
              </Typography>
              <br />
              <Grid container justify="center" >
                <FormControl required className={classes.formControl}>
                  <InputLabel id="demo-simple-select-required-label">เครื่อง (Video on demand Client)</InputLabel>
                  <Select
                    labelId="demo-simple-select-required-label"
                    id="client"
                    value={clientID}
                    className={classes.selectEmpty}
                    onChange={ClientIDhandleChange}
                    style={{ width: 300 }}
                  >
                    {clients.filter((filter: any) => filter.edges.state.sTATUSNAME == "Available").map((item: EntClientEntity) => (
                      <MenuItem value={item.id}>{item.cLIENTNAME}</MenuItem>
                    ))}
                  </Select>
                  <FormHelperText>Required</FormHelperText>
                </FormControl>
                <FormControl required className={classes.formControl}>
                  <InputLabel id="demo-simple-select-required-label">สมาชิกห้องสมุด</InputLabel>
                  <Select
                    disabled={true}
                    labelId="demo-simple-select-required-label"
                    id="library_member"
                    value={userID}
                    className={classes.selectEmpty}
                    style={{ width: 300 }}
                  >
                    {users.map((item: EntUser) => (
                      <MenuItem value={item.id}>{item.uSEREMAIL}</MenuItem>
                    ))}
                  </Select>

                </FormControl>
                <br />
                <FormControl required className={classes.formControl}>
                  <InputLabel id="demo-simple-select-required-label">จุดบริการ</InputLabel>
                  <Select
                    labelId="demo-simple-select-required-label"
                    id="service_point"
                    value={servicepointID}
                    className={classes.selectEmpty}
                    onChange={ServicePointIDhandleChange}
                    style={{ width: 300 }}
                  >
                    {servicepoint.map((item: EntServicePoint) => (
                      <MenuItem value={item.id}>{item.cOUNTERNUMBER}</MenuItem>
                    ))}
                  </Select>
                  <FormHelperText>Required</FormHelperText>
                </FormControl>
                <FormControl required className={classes.formControl}>
                  <TextField style={{ width: 300 }} error={userNumberError ? true : false} id="user_number"
                    helperText={userNumberError} type="number" InputProps={{ inputProps: { min: 1, max: 6 } }}
                    onChange={handleChange} label="จำนวนผู้ใช้บริการ (รวมผู้จอง)*"
                    value={bookingField.user_number} />


                  <br />
                </FormControl>
                <FormControl required className={classes.formControl}>
                  <TextField style={{ width: 300 }} error={borrowItemError ? true : false} id="borrow_item"
                    helperText={borrowItemError} type="number" InputProps={{ inputProps: { min: 1, max: 6 } }}
                    onChange={handleChange} label="HeadSet ที่ยืมร่วมกับการใช้เครื่อง (ชิ้น)*"
                    value={bookingField.borrow_item} />
                </FormControl>
                <FormControl required className={classes.formControl}>
                  <TextField style={{ width: 300 }} error={phoneNumberError ? true : false} id="phone_number"
                    helperText={phoneNumberError} onChange={handleChange} label="หมายเลขโทรศัพท์*"
                    value={bookingField.phone_number || ''} />
                </FormControl>

                <Grid container justify="center" item xs={12}>
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
                </Grid>
                <Grid container justify="center" item xs={12}>


                  {status ? (
                    <div>
                      {alert ? (
                        <Alert severity="success" style={{ width: 400 }} onClose={() => { setStatus(false); window.location.reload(false); }}>
                          <AlertTitle>Success</AlertTitle>
                          <div>
                            บันทึกข้อมูลสำเร็จ — <strong>🎉</strong>
                          </div>
                          <br />
                        </Alert>
                      ) : (
                          <Alert severity="error" style={{ width: 400 }} onClose={() => { setStatus(false); window.location.reload(false); }}>
                            <AlertTitle>Error</AlertTitle >
                            <div>
                              {alertMessage} — <strong>❌</strong>
                            </div><br />
                          </Alert>
                        )}
                    </div>
                  ) : null}
                </Grid>
              </Grid>
            </InfoCard>
          </Grid>
          <Grid item xs={12} md={5}>
            <InfoCard title="ตารางเครื่องรับชม VideoOnDemand">
              <ComponanceTable></ComponanceTable>
            </InfoCard>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
}