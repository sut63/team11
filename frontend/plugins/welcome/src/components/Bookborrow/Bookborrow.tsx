import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import TextField from '@material-ui/core/TextField';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import { EntUser } from '../../api/models/EntUser';
import { EntBook } from '../../api/models/EntBook';
import { EntServicePoint } from '../../api/models/EntServicePoint';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';

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
      width: '25ch',
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
    formControl: {
      width: 400,
    },
    selectEmpty: {
      marginTop: theme.spacing(2),
    },
    button: {
      margin: theme.spacing(1),
    },
  }),
);

export default function Create() {
  const name = JSON.parse(String(localStorage.getItem("userName")));
  const userName = "ยินดีต้อนรับ " + name
  const classes = useStyles();
  const api = new DefaultApi();

  const [users, setUsers] = useState<EntUser[]>([]);
  const [books, setBooks] = useState<EntBook[]>([]);
  const [servicepoints, setServicepoints] = useState<EntServicePoint[]>([]);

  const [userID, setUser] = useState(Number);
  const [bookID, setBook] = useState(Number);
  const [servicePointID, setServicepoint] = useState(Number);

  const [status, setStatus] = useState(false);
  const [loading, setLoading] = useState(true);

  const [dayOfBorrow, setDayOfBorrow] = useState(Number);
  const [pickup, setPickup] = useState(String);
  const [phoneNumber, setPhoneNumber] = useState(String);

  const idString = JSON.parse(String(localStorage.getItem("userID")));
  const idInt = parseInt(idString);

  const [dayOfBorrowerror, setDayOfBorrowerror] = React.useState('');
  const [pickuperror, setPickuperror] = React.useState('');
  const [phoneNumbererror, setPhoneNumbererror] = React.useState('');
  const [errormessege, setErrorMessege] = useState(String);
  const [alerttype, setAlertType] = useState(String);

  useEffect(() => {
    const getBook = async () => {
      const res = await api.listBookfrees({ limit: 10, offset: 0 });
      setLoading(false);
      setBooks(res);
    };
    getBook();

    const getServicepoint = async () => {
      const res = await api.listServicepoint({ limit: 10, offset: 0 });
      setLoading(false);
      setServicepoints(res);
    };
    getServicepoint();

    const getUser = async () => {
      const res = await api.listUser();
      setLoading(false);
      setUsers(res);
    };
    getUser();

    setUser(idInt);

  }, [loading]);

  //----------------

  const validateDayofborrow = (val: number) => {
    return val <= 14 && val >= 1 ? true : false

  }

  const validatePickup = (val: string) => {
    return val.length > 0  ? true : false
  }

  const validatePhonenumber = (val: string) => {
    return val.match("[0]\\d{9}");
  }

  const checkPattern = (id: string, value: string) => {
    console.log(value);
    switch (id) {
      case 'phoneNumber':
        validatePhonenumber(value) ? setPhoneNumbererror('') : setPhoneNumbererror('หมายเลขโทรศัพท์ขึ้นต้นด้วย 0 และตามด้วยตัวเลขอีก 9 ตัว');
        return;
      case 'pickup':
        validatePickup(value) ? setPickuperror('') : setPickuperror('กรุณากรอกข้อมูล ไม่สามารถเป็นค่าว่างได้');
        return;
      case 'dayOfBorrow':
        validateDayofborrow(Number(value)) ? setDayOfBorrowerror('') : setDayOfBorrowerror('จำนวนวันที่ยืมได้คือ 1-14 วัน');
        return;
      default:
        return;
    }
  }

  const handleBookchange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setBook(event.target.value as number);
  };

  const handleServicepointchange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setServicepoint(event.target.value as number);
  };

  const handlePhoneNumberChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkPattern('phoneNumber', validateValue)
    setPhoneNumber(event.target.value as string);
  };

  const handlePickupChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkPattern('pickup', validateValue)
    setPickup(event.target.value as string);
  };

  const handleDayOfBorrowChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkPattern('dayOfBorrow', validateValue)
    setDayOfBorrow(event.target.value as number);
  };

  const listbb = () => {
    setStatus(false);
    if(errormessege == "บันทึกข้อมูลสำเร็จ"){
    window.location.href ="http://localhost:3000/Bookborrow";
    }
  };

  const checkCaseSaveError = (field: string) => {
    if (field == "DAY_OF_BORROW") { setErrorMessege("ข้อมูลfield จำนวนวันที่ยืมผิด"); }
    else if (field == "PICKUP") { setErrorMessege("ข้อมูลfield วิธีรับหนังสือผิด"); }
    else if (field == "PHONE_NUMBER") { setErrorMessege("ข้อมูลfield เบอร์โทรศัพท์ผิด"); }
    else { setErrorMessege("บันทึกไม่สำเร็จใส่ข้อมูลไม่ครบ"); }
  }

  const createBookBorrow = async () => {
    if ((userID != 0) && (bookID != 0) && (servicePointID != 0)) {
      const apiUrl = 'http://localhost:8080/api/v1/bookborrows';
      const bookBorrow = {
        bookID: bookID,
        servicePointID: servicePointID,
        userID: userID,
        dayOfBorrow: Number(dayOfBorrow),
        pickup: pickup,
        phoneNumber: phoneNumber,
      };
      console.log(bookBorrow);
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(bookBorrow),
      };

      fetch(apiUrl, requestOptions)
        .then(response => response.json())
        .then(data => {
          console.log(data);
          setStatus(true);
          if (data.status === true) {
            setErrorMessege("บันทึกข้อมูลสำเร็จ");
            setAlertType("success");
          }
          else {
            checkCaseSaveError(data.error.Name);
            setAlertType("error");
          }
        });
    }
    else {
      setErrorMessege("บันทึกไม่สำเร็จใส่ข้อมูลไม่ครบ");
      setAlertType("error");
      setStatus(true);
    }
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
        title={`Welcome to BookBorrow System`}
        subtitle="ระบบยืมหนังสือ"
      >
        <Button
          // disabled={LogoutBtn}
          variant="contained"
          color="secondary"
          className={classes.button}
          startIcon={<LockOutlinedIcon />}
          onClick={() => {
            resetLocalStorage();
          }}>
          ล็อกเอ้าท์
          </Button>
      </Header>

      <Content>
        <ContentHeader title={userName}>
          {status ? (
            <div>
              {alerttype != "" ? (
                <Alert severity={alerttype} onClose={() => { listbb() }}>
                  {errormessege}
                </Alert>
              ) : null}
            </div>
          ) : null}

        </ContentHeader>

        <div className={classes.root}>
          <form noValidate autoComplete="off">

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div className={classes.paper}><strong>หนังสือ</strong></div>
                <InputLabel id="book-label"></InputLabel>
                <Select
                  labelId="ชื่อหนังสือ"
                  id="bookID"
                  value={bookID}
                  onChange={handleBookchange}
                  style={{ width: 400 }}
                >
                  {books.map((item: EntBook) => (
                    <MenuItem value={item.id}>{item.bookName}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </div>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div className={classes.paper}><strong>ผู้ยืม</strong></div>
                <InputLabel id="user-label"></InputLabel>
                <Select
                  disabled={true}
                  labelId="ผู้ยืม"
                  id="userID"
                  value={userID}
                  className={classes.selectEmpty}
                  style={{ width: 400 }}
                >
                  {users.map((item: EntUser) => (
                    <MenuItem value={item.id}>{item.uSERNAME}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </div>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div className={classes.paper}><strong>จุดบริการ</strong></div>
                <InputLabel id="servicepoint-label"></InputLabel>
                <Select
                  labelId="จุดบริการที่ยืมหนังสือ"
                  id="servicepointID"
                  value={servicePointID}
                  onChange={handleServicepointchange}
                  style={{ width: 400 }}
                >
                  {servicepoints.map((item: EntServicePoint) => (
                    <MenuItem value={item.id}>{item.cOUNTERNUMBER}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </div>

            <div className={classes.paper}><strong>จำนวนวันที่ยืม</strong></div>
            <TextField 
              style={{ width: 400, marginLeft: 20, marginRight: -10 }}
              id="dayOfBorrow"
              error = {dayOfBorrowerror ? true : false}
              label=""
              variant="outlined"
              color="primary"
              type="string"
              size="medium"
              helperText= {dayOfBorrowerror}
              value={dayOfBorrow}
              onChange={handleDayOfBorrowChange}
            />

            <div className={classes.paper}><strong>วิธีรับหนังสือ</strong></div>
            <TextField 
              style={{ width: 400, marginLeft: 20, marginRight: -10 }}
              id="pickup"
              error = {pickuperror ? true : false}
              label=""
              variant="outlined"
              color="primary"
              type="string"
              size="medium"
              helperText= {pickuperror}
              value={pickup}
              onChange={handlePickupChange}
            />

            <div className={classes.paper}><strong>เบอร์โทรศัพท์</strong></div>
            <TextField 
              style={{ width: 400, marginLeft: 20, marginRight: -10 }}
              id="phoneNumber"
              error = {phoneNumbererror ? true : false}
              label=""
              variant="outlined"
              color="primary"
              type="string"
              size="medium"
              helperText= {phoneNumbererror}
              value={phoneNumber}
              onChange={handlePhoneNumberChange}
            />

            <div className={classes.margin}>
              <Button
                onClick={() => {
                  createBookBorrow();
                }}
                variant="contained"
                color="primary"
              >
                ยืนยันการบันทึก
             </Button>
              <Button
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/"
                variant="contained"
              >
                ย้อนกลับ
             </Button>
            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
}

