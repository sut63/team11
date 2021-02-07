import React, { useState, useEffect } from 'react';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Book from '@material-ui/icons/MenuBook';
import VideoOnDemand from '@material-ui/icons/DesktopWindows';
import Room from '@material-ui/icons/EventSeat';
import Research from '@material-ui/icons/NoteAdd';
import Borrow from '@material-ui/icons/AddShoppingCart';
import Return from '@material-ui/icons/AssignmentReturn';
import { Link as RouterLink } from 'react-router-dom';
import {
  Typography,
  Grid,
  makeStyles,
  Button,
} from '@material-ui/core';
import {
  Content,
  InfoCard,
  Header,
  HomepageTimer,
  Page,
  pageTheme,
  ContentHeader,
  WarningPanel,
} from '@backstage/core';

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
    margin: theme.spacing(3),
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
    flexGrow: 1,
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
}));
const WelcomePage = () => {
  const [userName, setName] = useState("Getting Started \"Please Login\"");
  const [LibrarianBtn, setLibrarianBtn] = useState(Boolean);
  const [MemberBtn, setMemberBtn] = useState(Boolean);
  const [LoginBtn, setLoginBtn] = useState(Boolean);
  const [LogoutBtn, setLogoutBtn] = useState(Boolean);
  const classes = useStyles();
  
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const role = JSON.parse(String(localStorage.getItem("role")));
    const name = JSON.parse(String(localStorage.getItem("userName")));
    setLoginBtn(false);
    setLogoutBtn(true);
    if(role === "Librarian"){
      setLoading(false);
      setLibrarianBtn(false);
      setMemberBtn(true);
      setName("ยินดีต้อนรับ "+name)
      setLoginBtn(true);
      setLogoutBtn(false);
      
    }
    else if (role === "Library Member"){
      setLoading(false);
      setLibrarianBtn(true);
      setMemberBtn(false);
      setName("ยินดีต้อนรับ "+name)
      setLoginBtn(true);
      setLogoutBtn(false);
    }
    else{
      setLibrarianBtn(true);
      setMemberBtn(true);
    }

  }, [loading]);
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
        title={'Welcome to Library System'}
        subtitle="The Center for Library Resources and Educational Media"
      >
        <HomepageTimer />
      </Header>
      <Content>
        <ContentHeader title={userName}>
          <Button
            disabled={LoginBtn}
            variant="contained"
            color="primary"
            className={classes.button}
            startIcon={<LockOutlinedIcon />}
            component={RouterLink}
            to="/SignIn">
            ล็อกอิน
          </Button>
          <Button
            disabled={LogoutBtn}
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

        <Grid container justify="center">
          <Grid item xs={12}>
            <WarningPanel
              title="ประกาศ"
              message={
                <>
                  เวลาเปิดทำการ
                  วันจันทร์ - วันศุกร์ เวลา 08.30 - 16.30 น.
                  ปิดบริการเสาร์ อาทิตย์ และตามประกาศของมหาวิทยาลัย
                </>
              }
            />
          </Grid>
          <Grid item xs={12} md={7}>

            <InfoCard title="ห้องสมุด 🔎">
              <Typography variant="body1" gutterBottom>
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;ห้องสมุด 📚 คือแหล่งสารนิเทศ บริการทรัพยากรสารนิเทศในรูปแบบต่าง ๆ เช่น หนังสือ วารสาร หนังสือพิมพ์ จุลสาร กฤตภาค วัสดุเทป
                และโทรทัศน์ ซีดีรอม วีซีดี ดีวีดี โดยมีบรรณารักษ์ เป็นผู้ดำเนินงาน และบริหารงานต่าง ๆ ในห้องสมุด โดยจัดระบบเป็นหมวดหมู่
                และระเบียบเรียบร้อย เพื่อให้ผู้ใช้ห้องสมุดมีความสะดวกสืบค้นได้ง่ายและตรงกับความต้องการ
              </Typography>
              <Typography variant="h6" gutterBottom>
                ห้องสมุดในปัจจุบัน 💻
              </Typography>
              <Typography variant="body1" paragraph>
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;ทำหน้าที่เก็บรวบรวม จัดระบบ เพื่อให้บริการสื่อสารนิเทศต่าง ๆ ตลอดจนถึงเทคโนโลยีทางคอมพิวเตอร์ และเทคโนโลยีทางการสื่อสาร อีกทั้งยังมีเครื่องมือในการค้นหาและดำเนินการให้บริการสื่อต่าง ๆ เกิดประโยชน์สูงสุดแก่ผู้ใช้ห้องสมุด
              </Typography>
              <Typography variant="h6" gutterBottom>
                บทบาทของห้องสมุด 🏫
              </Typography>
              <Typography variant="body1" paragraph>
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;ห้องสมุด มีบทบาทต่อบุคคลต่าง ๆ มากมาย ซึ่งห้องสมุด สามารถทำประโยชน์ต่อสังคมในด้านต่าง ๆ <br /><br />
              &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;ด้านวัฒนธรรม ห้องสมุดเป็นที่บำรุงรักษาวัฒนธรรมของชาติหนังสือ ให้สืบทอดไปยังอนุชนรุ่นต่อไป เนื่องจากห้องสมุด เป็นแหล่งที่จัดเก็บข้อมูลทางด้านสารนิเทศ ซึ่งส่งเสริมศิลปะและวัฒนธรรมต่างๆ และสามารถใช้บ่งบอกความเจริญก้าวหน้าของประเทศนั้นๆ อีกด้วย
              วัตถุประสงค์หลัก ทั่วไปของห้องสมุดมีอยู่ด้วยกัน 5 ประการ ดังนี้<br /><br />
              &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
              1. เพื่อการศึกษา - การใช้ห้องสมุด เป็นหัวใจของการศึกษา เพราะห้องสมุดเป็นแหล่งรวบรวมความรู้ สามารถให้ผู้คน รู้จักศึกษาหาความรู้ด้วยตนเอง เพื่อให้เกิดความรู้เพิ่มเติมจากที่ร่ำเรียน ไม่ว่าจะเป็น นักเรียนอนุบาล ประถมศึกษามัธยมศึกษา อุดมศึกษา รวมไปถึงผู้ที่จบการศึกษาแล้ว ประชาชนทั่วไป สามารถใช้ห้องสมุดศึกษาหาความรู้ได้ตลอดชีวิต <br /><br />
              &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
              2. เพื่อให้ความรู้และข่าวสาร - ทุกวันนี้วิทยาการต่างๆ ก้าวไปอย่างรวดเร็ว ห้องสมุดเป็นสถานที่สำหรับศึกษาวิทยาการต่างๆ และติดตามข่าวความเคลื่อนไหวทั้งภายในและภายนอกประเทศทั่วโลก ให้คนรู้จักข่าวคราว ทันสมัย ทันต่อเหตุการณ์ มีความคิดริเริ่มสร้างสรรค์ ช่วยพัฒนาประเทศต่อไป<br /><br />
              &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
              3. เพื่อใช้ในการค้นคว้า - ในห้องสมุด มีสารนิเทศมากมาย เพื่อให้บริการแก่ผู้ใช้ เหมาะสำหรับเป็นเครื่องมือในการวิจัย ผู้ที่จะทำการวิจัย จำเป็นต้องใช้ข้อมูลเบื้องต้นจากเรื่องที่มีอยู่แล้ว<br /><br />
              &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
              4. เพื่อจรรโลงใจมนุษย์ - การอ่านหนังสือ นอกเสียจากการได้ความรู้แล้ว ยังทำให้ผู้อ่านมีความสุขได้อีกด้วย เนื่องจากความซาบซึ้งในความคิดที่ดีงาม ให้ความจรรโลงใจในสิ่งที่ดีแก่ผู้อ่าน ก่อให้เกิดแรงบันดาลใจ ให้ทำประโยชน์ต่อสังคม เช่น การอ่านหนังสือธรรมะ ทำให้ซาบซึ้งถึงคำสั่งสอนของพระพุทธเจ้า หนังสือวรรณกรรม วรรณคดี ทำให้เกิดจินตนาการ
              หนังสือประเภทต่างๆ เพื่อให้ความบันเทิงเบาสมอง เช่น นวนิยาย เรื่องสั้น นิตยสาร เป็นต้น ทำให้ผู้ใช้ ใช้เวลาว่างให้เกิดประโยชน์ และปลูกสำนึกรักการอ่าน รู้จักอ่านหนังสืออย่างมีวิจารณญาณ<br /><br />
              &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
              5. เพื่อนันทนาการ - หรือพักผ่อนหย่อนใจ ซึ่งเป็นสิ่งจำเป็นสำหรับชีวิตมนุษย์ที่ต้องการพักผ่อนสมอง ผ่อนคลายความตึงเครียด เกิดความเพลิดเพลินใจ สบายใจ วัสดุห้องสมุดที่มีเนื้อหาสาระ ทั้งสารคดีและบันเทิงคดี สามารถก่อให้เกิดนันทนาการได้<br /><br />
              </Typography>
            </InfoCard>
          </Grid>
          <Grid item xs={12} md={3}>
              <Grid>
            <InfoCard title="เมนูสำหรับการสร้าง 🛠">
                  <Button
                    disabled ={LibrarianBtn}
                    variant="contained"
                    color="secondary"
                    className={classes.button}
                    startIcon={<Research />}
                    component={RouterLink}
                    to="/Research"
                  >
                    เพิ่มงานวิจัย
                  </Button>
                  <br />
                  <Button
                    disabled ={LibrarianBtn}
                    variant="contained"
                    color="secondary"
                    className={classes.button}
                    startIcon={<Book />}
                    component={RouterLink}
                    to="/Book"
                  >
                    เพิ่มหนังสือ
                  </Button>
                  <br/>
                  <Button
                    disabled ={MemberBtn}
                    variant="contained"
                    color="secondary"
                    className={classes.button}
                    startIcon={<VideoOnDemand />}
                    component={RouterLink}
                    to="/VideoOnDemand"
                  >
                    จองเครื่อง VOD
                  </Button>
                  <br />
                  <Button
                    disabled ={MemberBtn}
                    variant="contained"
                    color="secondary"
                    className={classes.button}
                    startIcon={<Room />}
                    component={RouterLink}
                    to="/Room"
                  >
                    จองห้องติว
                  </Button>
                  <br />

                  <Button
                    disabled ={MemberBtn}
                    variant="contained"
                    color="secondary"
                    className={classes.button}
                    startIcon={<Borrow />}
                    component={RouterLink}
                    to="/Bookborrow"
                  >
                    ยืมหนังสือ
                  </Button>
                  <br />
                  <Button
                    disabled ={MemberBtn}
                    variant="contained"
                    color="secondary"
                    className={classes.button}
                    startIcon={<Return />}
                    component={RouterLink}
                    to="/Bookreturn"
                  >
                    คืนหนังสือ
                  </Button>
            </InfoCard>
            </Grid>
            <br/>
            <Grid>
            <InfoCard title="เมนูสำหรับการค้นหา 🔎">
                  <Button
                    disabled ={MemberBtn}
                    variant="contained"
                    color="primary"
                    className={classes.button}
                    startIcon={<Research />}
                    component={RouterLink}
                    to="/SearchResearch"
                  >
                    ค้นหางานวิจัย
                  </Button>
                  <br />
                  <Button
                    disabled ={MemberBtn}
                    variant="contained"
                    color="primary"
                    className={classes.button}
                    startIcon={<Book />}
                    component={RouterLink}
                    to="/SearchBook"
                  >
                    ค้นหาข้อมูลหนังสือ
                  </Button>
                  <br/>
                  <Button
                    disabled ={MemberBtn}
                    variant="contained"
                    color="primary"
                    className={classes.button}
                    startIcon={<VideoOnDemand />}
                    component={RouterLink}
                    to="/SearchBooking"
                  >
                    ค้นหาการจองเครื่องรับชม Video On Demand
                  </Button>
                  <br />
                  <Button
                    
                    variant="contained"
                    color="primary"
                    className={classes.button}
                    startIcon={<Room />}
                    component={RouterLink}
                    to="/SearchRoom"
                  >
                    ค้นหาใบจอง
                  </Button>
                  <br />

                  <Button
                    disabled ={MemberBtn}
                    variant="contained"
                    color="primary"
                    className={classes.button}
                    startIcon={<Borrow />}
                    component={RouterLink}
                    to="/SearchBookborrow"
                  >
                    ค้นหาการยืมหนังสือ
                  </Button>
                  <br />
                  <Button
                    disabled ={MemberBtn}
                    variant="contained"
                    color="primary"
                    className={classes.button}
                    startIcon={<Return />}
                    component={RouterLink}
                    to="/SearchBookreturn"
                  >
                    คืนหนังสือ
                  </Button>
            </InfoCard>
            </Grid>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
};

export default WelcomePage;
