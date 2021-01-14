
import React, { useState, useEffect } from 'react';
import 'date-fns';
import AssignmentIcon from '@material-ui/icons/Assignment';
import { Link as RouterLink } from 'react-router-dom';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormHelperText from '@material-ui/core/FormHelperText';
import IconButton from '@material-ui/core/IconButton';
import Select from '@material-ui/core/Select';
import AccountCircle from '@material-ui/icons/AccountCircle';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import MenuIcon from '@material-ui/icons/Menu';

import Grid from '@material-ui/core/Grid';

import {
  MuiPickersUtilsProvider,
  KeyboardTimePicker,
  KeyboardDatePicker,
} from '@material-ui/pickers';

import { EntRoominfo } from '../../api/models/EntRoominfo'; // import interface Roominfo
import { EntUser } from '../../api/models/EntUser'; // import interface User
import { EntPreemption } from '../../api/models/EntPreemption'; // import interface Preemption
import { EntPurpose } from '../../api/models/EntPurpose'; // import interface Purpose
import ComponanceTable from '../RoomTable';
import Swal from 'sweetalert2';

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
   container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  table: {
    minWidth: 650,
  },
  large: {
    width: theme.spacing(7),
    height: theme.spacing(7),
    
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
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const api = new DefaultApi();
  const [roominfos, setRoominfos] = React.useState<EntRoominfo[]>([]);
  const [preemptions, setPreemptions] = React.useState<EntPreemption[]>([]);
  const [purposes, setPurposes] = React.useState<EntPurpose[]>([]);
  
  const [loading, setLoading] = useState(true);
  
  const [userid, setuser] = useState(Number);
  const [roomid, setroom] = useState(Number);
  const [purposeid, setpurpose] = useState(Number);
  const [preempttime, setpreempttime] = useState(String);
  useEffect(() => {
    
    
 
  const getRoominfos = async () => {
      const res = await api.listRoominfo({ limit: 10, offset: 0 });
      setRoominfos(res);
    };
  getRoominfos();

  const getPurposes = async () => {
    const res = await api.listPurpose({ limit: 10, offset: 0 });
    setPurposes(res);
    
  };
  getPurposes();

  

}, [loading]);

const getPreemptions = async () => {
  const res = await api.listPreemption({ limit: 10, offset: 0 });
  setPreemptions(res);
  
};

const RoomhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setroom(event.target.value as number);
};

const PurposehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setpurpose(event.target.value as number);
};

const TimehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setpreempttime(event.target.value as string);
};

/*const Createpreempt = async () => { 
const preemption = {
  user: 1,
  roominfo: roomid,
  purpose: purposeid,

  added: preempttime + ":00+07:00"

};
const roon = {
  
  roomStatus: "ไม่ว่าง",
  

};
console.log(preemption);
console.log(roomid);
const res: any = await api.createPreemption({ preemption: preemption });
const ros = await api.updateRoominfo({id:roomid,user:roon});
setStatus(true);
   if (res.id != ''){
     setAlert(true);
   } else {
     setAlert(false);
   }
   
 };
*/
const Createpreempt = async () => {
  const preemption = {
    user: 1,
    roominfo: roomid,
    purpose: purposeid,
  
    added: preempttime + ":00+07:00"
  
  };
  const roon = {
    
    roomStatus: "ไม่ว่าง",
    
  
  };
  
  const apiUrl = 'http://localhost:8080/api/v1/preemptions'; 
  const requestOptions = { 
    method: 'POST', 
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(preemption),
};
  const ros = await api.updateRoominfo({id:roomid,user:roon});
  console.log(preemption); 
  fetch(apiUrl, requestOptions) 
    .then(response => response.json())
    .then(data => {
      console.log(data);
      if (data.status === true) { 
        
        Toast.fire({
          icon: 'success',
          title: 'จองสำเร็จ',
        });
      } else {
        Toast.fire({
          icon: 'error',
          title: 'จองไม่สำเร็จกรอกข้อมูลให้ครบ',
        });
      }
    });
}


 

 return (
   <Page theme={pageTheme.home}>
      
     
     
     <Header
       title={`จองห้องค้นคว้า `}
       subtitle=""
       
     >
       

       <Button variant="contained" color="secondary">
        log out
      </Button>
      <div className={classes.margin}>
      <Link component={RouterLink} to="/user">
      <IconButton>
                
                
   
               <AssignmentIcon fontSize="large" /></IconButton></Link>
               
      </div>                    
            <div>
              <IconButton>
                
                
              
                <AccountCircle className={classes.large}  />
                
              </IconButton>
              
            </div>
         
        
      
     </Header>
     <Content>

       
       <ContentHeader title="">
       

         
         
       </ContentHeader>
       <div className={classes.root}>
         <form noValidate autoComplete="off">
           
         <ComponanceTable></ComponanceTable>


    <div className={classes.margin}>
             
             </div>
    



         <FormControl className={classes.root}>
         <InputLabel id="se">เลือกห้อง และ เวลา</InputLabel>
         <Select
                  name="room"
                  value={roomid}
                  onChange={RoomhandleChange}
                >
                  {roominfos.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.roomNo+" "+item.roomTime}
                        
                      </MenuItem>
                    );
                  })}
                </Select>
      </FormControl>

      <div className={classes.margin}>
             
             </div>

      <FormControl className={classes.root}>
         <InputLabel id="pur">วัตถุประสงค์</InputLabel>
          <Select
         name="purpose"
          value={purposeid}
          onChange={PurposehandleChange}
        >
          {purposes.map(item => {
            return (
              <MenuItem key={item.id} value={item.id}>
                {item.purposeName}
                
              </MenuItem>
            );
          })}
        </Select>
      </FormControl>

      <div className={classes.margin}>
             
             </div>

      <form className={classes.container} noValidate>
      <TextField
        id="datetime-local"
        label="เวลาที่จอง"
        type="datetime-local"
        
        value={preempttime}
        onChange={TimehandleChange}
        className={classes.textField}
        InputLabelProps={{
          shrink: true,
        }}
      />
    </form>
    {status ? (
           <div>
             {alert ? (
               <Alert severity="success" style={{ marginTop: 20 }}>
                 จองสำเร็จ
                 
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 Aww จองบ่ได๋
               </Alert>
             )}
           </div>
         ) : null}      
    <div className={classes.margin}>
             
             </div>
             
 
           <div className={classes.margin}>
             <Button
               onClick={() => {
                Createpreempt()
               }}
               variant="contained"
               color="primary"
             >
               จอง
             </Button>
             
           </div>
         </form>
       </div>
       
     </Content>
     
               
   </Page>
   
 );
}



