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
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import { EntAuthor } from '../../api/models/EntAuthor'; //-------
import { EntCategory } from '../../api/models/EntCategory'; //-------
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
interface addBook{
  
  bar_code?:  string;
  book_page?: number;
  book_name?: string;
}
export default function Create() {
  const name = JSON.parse(String(localStorage.getItem("userName")));
  const userName ="ยินดีต้อนรับ "+name

  const classes = useStyles();
  const api = new DefaultApi();
  //-------
  const [authors, setAuthors] = useState<EntAuthor[]>([]);
  const [categorys, setCategorys] = useState<EntCategory[]>([]);


  const [userid, setUser] = useState(Number);
  //-------
  const [authorid, setAuthor] = useState(Number);
  const [categoryid, setCategory] = useState(Number);


  const [loading, setLoading] = useState(true);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);


  const idString = JSON.parse(String(localStorage.getItem("userID")));
  const idInt = parseInt(idString);

  useEffect(() => {
    //-------
    const getAuthor = async () => {
      const res = await api.listAuthor({ offset: 0 });
      setLoading(false);
      setAuthors(res);
      console.log(res);
    };
    getAuthor();


    const getCategory = async () => {
      const res = await api.listCategory({ offset: 0 });
      setLoading(false);
      setCategorys(res);
      console.log(res);
    };
    getCategory();

    setUser(idInt);

  }, [loading]);


  //----------------
  const handleAuthorchange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setAuthor(event.target.value as number);
  };

  const handleCategorychange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setCategory(event.target.value as number);
  };
  const [bookNameError, setBookNameError] = useState('');
  const [barCodeError, setBarCodeError] = useState('');
  const [bookPageError, setBookPageError] = useState('');
  const [alertMessage, setAlertMessage] = useState('');
  const [addBook, setAddBook] = useState<Partial<addBook>>({});

  const handleChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as keyof typeof Create;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setAddBook({ ...addBook, [id]: value });
  };
  const validateBookName = (val: string) => {
    return val.length == 0 ? false : true;
  }
  const validateBookPage = (val: number) => {
    return val == 0 ? false : true
  }
  const validateBarCode = (val: string) => {
    return val.match("\\d{10}") && val.length == 10;
  }
  const checkPattern = (id: string, value: any) => {
    switch (id) {
      case 'book_name':
        validateBookName(value) ? setBookNameError('') : setBookNameError('โปรดใส่ชื่อหนังสือ');
        return;
      case 'book_page':
        validateBookPage(value) ? setBookPageError('') : setBookPageError('จำนวนหน้าหนังสือต้องมากกว่า 1 หน้า');
        return;
      case 'bar_code':
        validateBarCode(value) ? setBarCodeError('') : setBarCodeError('Barcode ต้องเป็นตัวเลขทั้งหมด 10 ตัว');
        return;
      default:
        return;
    }
  }
  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'BookName':
        setAlertMessage("error ข้อมูล field BookName ผิด");
        return;
      case 'Barcode':
        setAlertMessage("error ข้อมูล field Barcode ผิด");
        return;
      case 'BookPage':
        setAlertMessage("error ข้อมูล field BookPage ผิด");
        return;
      default:
        setAlertMessage("บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }
  const createBook = async () => {
      const book = {
        userid: userid,
        author: authorid,
        category: categoryid,
        bookname: addBook.book_name,
        bookPage: Number(addBook.book_page),
        barCode: addBook.bar_code

      };
      console.log(addBook)
      console.log(book)
      const apiUrl = 'http://localhost:8080/api/v1/books';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(book),
    };

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          setStatus(true);
          setAlert(true);
        } else {
          setStatus(true);
          setAlert(false);
          checkCaseSaveError(data.error.Name)
        }
      });
    const timer = setTimeout(() => {
      setStatus(false);
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
        title={`Welcome to Book System`}
        subtitle="กรอกข้อมูลหนังสือ"
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
              {alert ? (
                <Alert severity="success">
                  เพิ่มบันทึกเรียบร้อย!
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    {alertMessage}
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>

        <div className={classes.root}>
          <form noValidate autoComplete="off">

            <div className={classes.paper}><strong>ชื่อหนังสือ</strong></div>
            <TextField className={classes.textField}
              style={{ width: 400, marginLeft: 20, marginRight: -10 }}
              id="book_name"
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              error={bookNameError ? true : false}
              helperText={bookNameError}
              value={addBook.book_name || ''}
              onChange={handleChange}
            />
            <div className={classes.paper}><strong>จำนวนหน้าหนังสือ</strong></div>
            <TextField className={classes.textField}
              style={{ width: 400, marginLeft: 20, marginRight: -10 }}
              id="book_page"
              label=""
              variant="standard"
              color="secondary"
              size="medium"
              type="number" 
              InputProps={{ inputProps: { min: 1, max: 6 } }}
              error={bookPageError ? true : false}
              helperText={bookPageError}
              value={addBook.book_page}
              onChange={handleChange}
            />
            <div className={classes.paper}><strong>รหัส Barcode บนหนังสือ</strong></div>
            <TextField className={classes.textField}
              style={{ width: 400, marginLeft: 20, marginRight: -10 }}
              id="bar_code"
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              error={barCodeError ? true : false}
              helperText={barCodeError}
              value={addBook.bar_code}
              onChange={handleChange}
            />

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div className={classes.paper}><strong>ชื่อผู้แต่ง</strong></div>
                <InputLabel id="author-label"></InputLabel>
                <Select
                  labelId="author-label"
                  id="author"
                  value={authorid}
                  onChange={handleAuthorchange}
                  style={{ width: 400 }}
                >
                  {authors.map((item: EntAuthor) => (
                    <MenuItem value={item.id}>{item.name}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </div>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div className={classes.paper}><strong>ชื่อหมวดหนังสือ</strong></div>
                <InputLabel id="category-label"></InputLabel>
                <Select
                  labelId="category-label"
                  id="category"
                  value={categoryid}
                  onChange={handleCategorychange}
                  style={{ width: 400 }}
                >
                  {categorys.map((item: EntCategory) => (
                    <MenuItem value={item.id}>{item.categoryName}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </div>


            <div className={classes.margin}>
              <Button
                onClick={() => {
                  createBook();
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