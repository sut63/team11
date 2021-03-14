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

import { DefaultApi } from '../../api/apis';

import Swal from 'sweetalert2';

import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { EntPreemption } from '../../api/models/EntPreemption'; // import interface Preemption
import { EntUser } from '../../api/models/EntUser';// import interface User

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

   large: {
    width: theme.spacing(7),
    height: theme.spacing(7),
    
  },

  table: {
    minWidth: 650,
  },

 }),
);
 
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





 
export default function Preempttable() {
 const classes = useStyles();
 const [loading, setLoading] = useState(true);
 const api = new DefaultApi();
 const [preemptions, setPreemptions] = React.useState<EntPreemption[]>([]);
 
 const [users, setUsers] = useState<EntUser[]>(Array);
 const [name, setname] = useState(String);
 
 
 useEffect(() => {
  
    const getUsers = async () => {
      const res = await api.listUser();
      setLoading(false);
      setUsers(res);
    };
    getUsers();
    
}, [loading]);

const namehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setname(event.target.value as string);
    
    
}

const alertMessage = () =>{
  Toast.fire(
    {
      icon: 'success',
      title: '',
    }
  );
}


var lenreturn : number

const Seacrh = async () => {
  setPreemptions([]);
  users.map((item: any) => {  
    if (item.uSERNAME == name ) {
      const getPreemptions = async () => {
        const res = await api.listPreemption({nameu: name });
        setPreemptions(res); 
        console.log(res);   
        lenreturn = res.length
        if (lenreturn > 0){
          alertMessage()
      }else{  
        
      }  
      };
      getPreemptions();
    }   
   
    
  });

  if (lenreturn > 0){
    
}else{ 
  Toast.fire(
    {
      icon: 'error',
      title: 'ไม่พบข้อมูล',
    }
  );  
}   
    
}


 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={'ระบบค้นหาใบจอง'}
       subtitle=""
     >
         


      
     </Header>
     
     <Content>
      
       <ContentHeader title={""}>
         
       </ContentHeader>
       
       <div className={classes.margin}>
               
             </div>
       
       <div className={classes.root}>
         <form noValidate autoComplete="off">
         <FormControl className={classes.root}>
      <TextField 
      id="name" 
      label="กรอกชื่อ" 
      helperText="" 
      value={name}
      onChange={namehandleChange}
      />   

      </FormControl> 

      <div className={classes.margin}>
             <Button
               onClick={() => {
                Seacrh()
               }}
               
               variant="contained"
               color="primary"
             >
               ค้นหา
             </Button>
             
           </div>

           <div className={classes.margin}>
             
             </div>


         <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>ใบจองห้องค้นคว้า</TableCell>
            
            
           
          </TableRow>
        </TableHead>
        <TableBody>
        {preemptions.map(item => (
           <TableRow key={item.id}>
             <TableCell align="left">{"เลขใบจอง: "+item.id+" ห้อง: "+item.edges?.roomID?.roomNo+" เวลา:  "+item.edges?.roomID?.roomTime+" จองโดย: "+item.edges?.userID?.uSERNAME+" ผู้รับกุญแจ: "+item.surrogateid+" เบอร์โทรติดต่อ: "+item.surrogatephone}</TableCell>
             
             <TableCell align="right">
               
             </TableCell>
           </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
           
 
           
         </form>
       </div>
     </Content>
   </Page>
 );
}

