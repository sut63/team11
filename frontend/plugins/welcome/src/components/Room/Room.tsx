
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
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormHelperText from '@material-ui/core/FormHelperText';
import IconButton from '@material-ui/core/IconButton';
import Select from '@material-ui/core/Select';

import Grid from '@material-ui/core/Grid';
import Swal from 'sweetalert2';
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
  button: {
    margin: theme.spacing(1),
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
  const name = JSON.parse(String(localStorage.getItem("userName")));
  const userName ="ยินดีต้อนรับ "+name
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
  const [otherid, setotherid] = useState(String);
  const [phoneuser, setphoneuser] = useState(String);
  const [otherphone, setotherphone] = useState(String);
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

const PhoneuserhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setphoneuser(event.target.value as string);
};
const OtheridhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setotherid(event.target.value as string);
};
const OtherphonehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setotherphone(event.target.value as string);
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

//const res: any = await api.createPreemption({ preemption: preemption });
const ros = await api.updateRoominfo({id:roomid,user:roon});
console.log(preemption);
console.log(roomid);

        
        
        

setStatus(true);
   if ( res.roominfo = 400){
     setAlert(false);
   } else {
     setAlert(true);
   }
   const timer = setTimeout(() => {
    setStatus(false);
  }, 1000);
 }; */
const alertMessage = (icon: any, title: any) =>{
  Toast.fire(
    {
      icon: icon,
      title: title,
    }
  );
}

const checkCaseSaveError = (field: string) =>{
  switch(field){
    case 'Phonenumber':
      alertMessage("error","เบอร์โทรศัพท์ต้องมี10หลัก");
      return;
    case 'Surrogateid':
      alertMessage("error","รหัสนักศึกษาขึ้นต้อนด้วย B,M,D ตามด้วยตัวเลขเจ็ดหลัก");
      return;  
    case 'Surrogatephone':
      alertMessage("error","เบอร์โทรศัพท์ต้องมี10หลัก");
      return;
    default:
      alertMessage("error","กรุณากรอกข้อมูลให้ครบทุกส่วน");  
      return;
  }
}

const updateroom = async () => {
  const roon = {
    roomStatus: "ไม่ว่าง",
  };
  const ros = await api.updateRoominfo({id:roomid,user:roon});
}


 const Createpreempt = async () => {
  const preemption = {
    user: 1,
    roominfo: roomid,
    purpose: purposeid,
    otherpeopleid: otherid,
    otherpeoplephone: otherphone,
    phoneuser: phoneuser,
    //added: preempttime + ":00+07:00"
  
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
  
  console.log(preemption); 
  fetch(apiUrl, requestOptions) 
    .then(response => response.json())
    .then(data => {
      console.log(data);
      if (data.status === true) { 
        updateroom();
        Toast.fire({
          icon: 'success',
          title: 'จองสำเร็จ',
        });
      } else {
        checkCaseSaveError(data.error.Name)
      }
    });
}

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
       title={`จองห้องค้นคว้า `}
       subtitle=""
       
     >
       
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

       
      <div className={classes.margin}>
      <Link component={RouterLink} to="/user">
      <IconButton>
                
                
   
               <AssignmentIcon fontSize="large" /></IconButton></Link>
               
      </div>                    
            
         
        
      
     </Header>
     <Content>

       
       <ContentHeader title={userName}>
       

         
         
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
      <FormControl className={classes.root}>
      <TextField 
      id="Phoneuser" 
      label="เบอร์โทรศัพท์มือถือผู้จอง" 
      helperText="เบอร์มือถือ10หลัก" 
      value={phoneuser}
      onChange={PhoneuserhandleChange}
      />   

      </FormControl>

      <div className={classes.margin}>
             
             </div>
      <FormControl className={classes.root}>
      <TextField 
      id="Ohterpeopleid" 
      label="รหัสนักศึกษาผู้มารับกุญแจ" 
      helperText="BXXXXXXX" 
      value={otherid}
      onChange={OtheridhandleChange}
      />   

      </FormControl> 
      
      <div className={classes.margin}>
             
             </div>
      <FormControl className={classes.root}>
      <TextField 
      id="Phoneother" 
      label="เบอร์โทรศัพท์มือถือผู้รับกุญแจ" 
      helperText="เบอร์มือถือ10หลัก" 
      value={otherphone}
      onChange={OtherphonehandleChange}
      />   

      </FormControl>


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



