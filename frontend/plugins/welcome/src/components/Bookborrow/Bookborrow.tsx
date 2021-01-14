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
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import { EntUser } from '../../api/models/EntUser';
import { EntBook } from '../../api/models/EntBook'; 
import { EntServicePoint } from '../../api/models/EntServicePoint'; 


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
  }),
);

export default function Create() {
  const classes = useStyles();
  const api = new DefaultApi();

  const [users, setUsers] = useState<EntUser[]>([]);
  //-------
  const [books, setBooks] = useState<EntBook[]>([]);
  const [servicepoints, setServicepoints] = useState<EntServicePoint[]>([]);

  //-------
  const [userID, setUser] = useState(Number);
  const [bookID, setBook] = useState(Number);
  const [servicePointID, setServicepoint] = useState(Number);

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);


  useEffect(() => {
    const getBook = async () => {
      const res = await api.listBook({ limit: 10, offset: 0 });
      setLoading(false);
      setBooks(res);
      console.log(res);
    };
    getBook();


    const getServicepoint = async () => {
      const res = await api.listServicepoint({ limit: 10, offset: 0 });
      setLoading(false);
      setServicepoints(res);
      console.log(res);
    };
    getServicepoint();

    const getUser = async () => {
      const res = await api.listUser({ limit: 10, offset: 0 });
      setLoading(false);
      setUsers(res);
      console.log(res);
    };
    getUser();

  }, [loading]);

  //----------------
  const handleBookchange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setBook(event.target.value as number);
  };

  const handleServicepointchange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setServicepoint(event.target.value as number);
  };

  const handleUserchange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUser(event.target.value as number);
  };
  const createBookBorrow = async ()=>{
    if ((userID != "") && (bookID!= "") && (servicePointID!= "")) {  
      const bookBorrow = {
        bookID: bookID,
        servicePointID: servicePointID,
        userID: userID,
      };
      console.log(bookBorrow);
      const res: any = await api.createBookborrow({ bookborrow: bookBorrow });
      setStatus(true);
        if(res.id != ''){
            setAlert(true);
            window.location.reload(false);
        }
        }else{
            setAlert(false);
            setStatus(true);
        }
        const timer = setTimeout(() => {
            setStatus(false);
        }, 4000);
    }
  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`ระบบยืมหนังสือ`}>
        </Header>
        
      <Content>
        <ContentHeader title="เพิ่มข้อมูลการยืมหนังสือ">
        {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 เพิ่มบันทึกเรียบร้อย!
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 บันทึกไม่สำเร็จ กรุณากรอกข้อมูลใหม่
               </Alert>
             )}
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
                  labelId="ผู้ยืม"
                  id="userID"
                  value={userID}
                  onChange={handleUserchange}
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