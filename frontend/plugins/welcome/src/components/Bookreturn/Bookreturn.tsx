import React, { useState, useEffect } from 'react';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import { Link as RouterLink } from 'react-router-dom';
import Swal from 'sweetalert2';
import {
  Content, Header, Page, pageTheme, ContentHeader,
} from '@backstage/core';
import {
   TextField, Button, makeStyles,
  Theme, FormControl, InputLabel, MenuItem,
  Select, createStyles,
} from '@material-ui/core';


import { DefaultApi } from '../../api/apis';
import {
  EntUser,
  EntBookborrow,
  EntLocation,
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
    button: {
      margin: theme.spacing(1),
    },
    formControl: {
      margin: theme.spacing(1),
      minWidth: 120,
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
    textField: {
      width: '50ch',
    },
  }),
);

// alert setting
const Toast = Swal.mixin({
  toast: true,
  position: 'top-end',
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
  didOpen: toast => {
    toast.addEventListener('mouseenter', Swal.stopTimer);
    toast.addEventListener('mouseleave', Swal.resumeTimer);
  },
});

export default function Create() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);

  const [users, setUsers] = React.useState<EntUser[]>(Array);
  const [bookborrows, setBookborrows] = React.useState<EntBookborrow[]>([]);
  const [locations, setLocations] = React.useState<EntLocation[]>(Array);

  const idString = JSON.parse(String(localStorage.getItem("userID")));
  const idInt = parseInt(idString);

  const [damagedpoint, setDamagedpoint] = useState(Number);
  const [damagedpointname, setDamagedpointname] = useState(String);
  const [lost, setLost] = useState(String);

  const [errordamagedpoint, setErrorDamagedpoint] = useState(true);
  const [errordamagedpointname, setErrorDamagedpointName] = useState(true);
  const [errorlost, setErrorLost] = useState(true);



  // function validate DamagedPoint
  const ValidateDamagedPoint = (val: number) => {
    val < -1 || val > 10 ? setErrorDamagedpoint(false) : setErrorDamagedpoint(true);
  }

  // function validate DamagedPointName
  const ValidateDamagedPointName = (val: string) => {
    val.match("^[a-zA-Z, ]+$") ? setErrorDamagedpointName(true) : setErrorDamagedpointName(false);
  }

  // function validate Lost
  const ValidateLost = (lost: string) => {
    lost.match("^[a-zA-Z]+$") ? setErrorLost(true) : setErrorLost(false);
  }

  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }

  const checkCaseSaveError = (field: string) => {
    switch (field) {
      case 'DAMAGED_POINT':
        alertMessage("error", "ต้องเป็นตัวเลข 1-10 เท่านั้น");
        return;
      case 'DAMAGED_POINTNAME':
        alertMessage("error", "จุดที่เสียหายเป็นภาษาอังกฤษเท่านั้น เช่น TopFront,BottomBack");
        return;
      case 'LOST':
        alertMessage("error", "ถ้าหายให้พิมพ์ lost ถ้าไม่พิมพ์ no");
        return;
      default:
        alertMessage("error", "บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }


  useEffect(() => {
    const getUser = async () => {
      const res = await api.listUser();
      setLoading(false);
      setUsers(res);
      console.log("users => " + users);
    };
    getUser();

    setUserID(idInt);

    const getBookborrows = async () => {
      const res = await api.getBookborrowuser({ id: idInt });
      setBookborrows(res);
      console.log("bookborrows => " + bookborrows);
    };
    getBookborrows();

    const getLocations = async () => {
      const res = await api.listLocation({ limit: 10, offset: 0 });
      setLoading(false);
      setLocations(res);
    };
    getLocations();

  }, [loading]);

  
  const BookborrowIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setBookborrowID(event.target.value as number);
  };

  const LocationIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setLocationID(event.target.value as number);
  };

  const DamagedPointhandleChange = (event: React.ChangeEvent<{ value: any }>) => {
    setDamagedpoint(event.target.value as number);
    ValidateDamagedPoint(event.target.value as number);
  };

  const DamagedPointNamehandleChange = (event: React.ChangeEvent<{ value: any }>) => {
    setDamagedpointname(event.target.value as string);
    ValidateDamagedPointName(event.target.value as string);
  };

  const LosthandleChange = (event: React.ChangeEvent<{ value: any }>) => {
    setLost(event.target.value as string);
    ValidateLost(event.target.value as string);
  };

  const [userID, setUserID] = useState(Number);
  const [bookborrowID, setBookborrowID] = useState(Number);
  const [locationID, setLocationID] = useState(Number);


  const createBookreturn = async () => {
    const apiUrl = 'http://localhost:8080/api/v1/bookreturns';

    const bookreturn = {
      userID: userID,
      bookborrowID: bookborrowID,
      locationID: locationID,
      damagedPoint: Number(damagedpoint),
      damagedPointName: String(damagedpointname),
      lost: String(lost),

    };
    console.log(bookreturn);

    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(bookreturn),
    };

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          
          checkCaseSaveError(data.error.Name)
        }
      });
      const timer = setTimeout(() => {
      window.location.reload(false);
    }, 4000);
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
        title={"ระบบคืนหนังสือ"}
      >

      </Header>
      <Content>
        <ContentHeader title="เพิ่มข้อมูลการคืนหนังสือ">
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


        <div className={classes.root}>
          <form noValidate autoComplete="off">
            

             
                  <div>
                    <FormControl
                      required className={classes.formControl}
                    >
                      
                      <InputLabel id="user-label">สมาชิกห้องสมุด</InputLabel>
                      <Select
                        disabled={true}
                        labelId="user-label"
                        id="user"
                        value={userID}
                        //onChange={UserIDhandleChange}
                        style={{ width: 300, height: '7vh' }}
                      >
                        {users.map((item: EntUser) => (
                          <MenuItem value={item.id}>{item.uSERNAME}</MenuItem>
                        ))}
                      </Select>
                    </FormControl>
                  </div>
               

                <div>
                <FormControl
                  required className={classes.formControl}
                >
                  <InputLabel id="bookborrow-label">รายการยืมหนังสือ</InputLabel>
                  <Select
                    labelId="bookborrow-label"
                    id="bookborrow"
                    value={bookborrowID}
                    onChange={BookborrowIDhandleChange}
                    style={{ width: 300, height: '7vh' }}
                  >
                    {bookborrows.map(item => {
                      return (
                        <MenuItem value={item.id}>{item.edges?.book?.bookName}</MenuItem>
                      );
                    })}
                  </Select>
                </FormControl>
                </div>

                <div>
                  <FormControl
                    required className={classes.formControl}
                  >
                    
                    <InputLabel id="location-label">สถานที่คืนหนังสือ</InputLabel>
                    <Select
                      labelId="location-label"
                      id="location"
                      value={locationID}
                      onChange={LocationIDhandleChange}
                      style={{ width: 300, height: '7vh' }}
                    >
                      {locations.map((item: EntLocation) => (
                        <MenuItem value={item.id}>{item.lOCATIONNAME}</MenuItem>
                      ))}
                    </Select>
                  </FormControl>
                  </div>

              <div>
                <FormControl
                  required className={classes.formControl}
                >
                  <TextField
                    style={{ width: 300 }}
                    error={errordamagedpoint ? false : true}
                    id="DAMAGED_POINT"
                    label="ตำแหน่งที่ชำรุด"
                    variant="outlined"
                    type="number"
                    size="medium"
                    helperText={errordamagedpoint ? "" : "ต้องเป็นตัวเลข 0-10 เท่านั้น"}
                    InputProps={{ inputProps: { min: -1, max: 10 } }}
                    value={damagedpoint}
                    onChange={DamagedPointhandleChange}
                  />
                </FormControl>
              </div>

              <div>
                <FormControl
                  required className={classes.formControl}
                >
                  <TextField
                    style={{ width: 300 }}
                    error={errordamagedpointname ? false : true}
                    id="DAMAGED_POINTNAME"
                    label="ชื่อตำแหน่งที่ชำรุด"
                    variant="outlined"
                    type="string"
                    size="medium"
                    helperText={errordamagedpointname ? "" : "จุดที่เสียหายเป็นภาษาอังกฤษเท่านั้น เช่น TopFront,BottomBack"}
                    value={damagedpointname}
                    onChange={DamagedPointNamehandleChange}
                  />
                </FormControl>
              </div>

              <div>
                <FormControl
                 required className={classes.formControl}
                >
                  <TextField
                    style={{ width: 300 }}
                    error={errorlost ? false : true}
                    id="LOST"
                    label="กรณีหนังสือหาย"
                    variant="outlined"
                    type="string"
                    size="medium"
                    helperText={errorlost ? "" : "ถ้าหายให้พิมพ์ lost ถ้าไม่ พิมพ์ no"}
                    value={lost}
                    onChange={LosthandleChange}
                  />
                </FormControl>
              </div>

              <td></td><td>
                <div className={classes.margin}>
                  <Button
                    onClick={() => {
                      createBookreturn();
                    }}
                    variant="contained"
                    color="primary"
                  >
                    บันทึกการคืน
           </Button>

                  <Button

                    style={{ marginLeft: 20 }}
                    component={RouterLink}
                    to="/"
                    variant="contained"
                  >

                    กลับ
           </Button>


                </div></td>
            
          </form>
        </div>

      </Content>
    </Page>
  );
}