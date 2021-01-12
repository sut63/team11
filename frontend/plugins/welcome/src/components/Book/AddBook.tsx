import React, { useState,useEffect} from 'react';
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
//import Autocomplete from '@material-ui/lab/Autocomplete';
//import Typography from '@material-ui/core/Typography';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
//import InputAdornment from '@material-ui/core/InputAdornment';
//import { EntCarservice } from '../../api/models/EntCarservice';
import { EntUser } from '../../api/models/EntUser';
import { EntAuthor } from '../../api/models/EntAuthor'; //-------
import { EntCategory } from '../../api/models/EntCategory'; //-------
import { EntBook } from '../../api/models/EntBook'; //-------

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
 const profile = {thisName: 'ระบบเพิ่มข้อมูลหนังสือ' };
 const api = new DefaultApi();
 
 const [users, setUsers] = useState<EntUser[]>([]);
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

 const [title, setTitle] = useState(String);
 
 useEffect(() => {
    //-------
    const getAuthor = async () => {
      const res = await api.listAuthor({offset :0});
      setLoading(false);
      setAuthors(res);
      console.log(res);
    };
    getAuthor();


    const getCategory = async () => {
      const res = await api.listCategory({offset :0});
      setLoading(false);
      setCategorys(res);
      console.log(res);
    };
    getCategory();

   const getUser = async () => {
    const res = await api.listUser({offset :0});
    setLoading(false);
    setUsers(res);
   };
   getUser();

}, [loading]);

const handleTitleChange = (event: any) => {
  setTitle(event.target.value as string);
 };

 //----------------
 const handleAuthorchange = (event: React.ChangeEvent<{value: unknown}>) => {
  setAuthor(event.target.value as number);
};

const handleCategorychange = (event: React.ChangeEvent<{value: unknown}>) => {
  setCategory(event.target.value as number);
};

 const handleUserchange = (event: React.ChangeEvent<{value: unknown}>) => {
  setUser(event.target.value as number);
};


 const createBook = async ()=>{
  if ((title != null) && (title != "") && (authorid != null) && (categoryid != null) ) {
   const book ={
    userid: 1,
    author: authorid,
    category: categoryid,
    bookname: title,
   };
   console.log(book);
   const res: any = await api.createBook({book : book});
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
        }, 3000);
 };
 
 return (
    <Page theme={pageTheme.home}>
     <Header
       title={`${profile.thisName}`}
       subtitle="กรอกข้อมูลหนังสือ"
     ></Header>

     <Content>
       <ContentHeader title="ระบบเพิ่มข้อมูลหนังสือ">
       
         {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 เพิ่มบันทึกเรียบร้อย!
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 บันทึกไม่สำเร็จ
               </Alert>
             )}
           </div>
         ) : null}
       </ContentHeader>

       <div className={classes.root}>
          <form noValidate autoComplete="off">
          
            <div className={classes.paper}><strong>ชื่อหนังสือ</strong></div>
            <TextField className={classes.textField}
            style={{ width: 400 ,marginLeft:20,marginRight:-10}}
              id="BookName"
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={title}
              onChange={handleTitleChange}
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
                id="ผู้เเต่ง"
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
                id="หมวดหนังสือ"
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