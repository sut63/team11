import React, { FC } from 'react';
import { Typography, Grid } from '@material-ui/core';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';

const HeaderCustom = {
  minHeight: '50px',
};

const useStyles = makeStyles({
  root: {
    maxWidth: 345,
  },
});

export type ProfileProps = {
  name: string; 
  id: string;
  system: string;
};

export function CardTeam({ name, id, system }: ProfileProps) {
  const classes = useStyles();
  return (
    <Grid item xs={12} md={3}>
      <Card className={classes.root}>
        <CardActionArea>
          <CardMedia
            component="img"
            alt="นาย สมชาย ใจดี"
            height="140"
            image="../../image/account.jpg"
            title="นาย สมชาย ใจดี"
          />
          <CardContent>
            <Typography gutterBottom variant="h5" component="h2">
              {system}
            </Typography>
            <Typography gutterBottom variant="h5" component="h2">
              {id} {name}
            </Typography>
          </CardContent>
        </CardActionArea>
      </Card>
    </Grid>
  );
}

const WelcomePage: FC<{}> = () => {
  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ระบบ...`}></Header>
      <Content>
        
        <ContentHeader title="สมาชิกในกลุ่ม"></ContentHeader>
        
        <Grid container>
          <CardTeam name={"นางสาว มนัสชนก ศรีเครือดง"} id={"B6102647"} system={"ระบบยืมหนังสือ"}></CardTeam>
          <CardTeam name={"นางสาว ศศิธร เจริญศิริ"} id={"B6105419"} system={"ระบบเพิ่มข้อมูลหนังสือ"}></CardTeam>
          <CardTeam name={"นาย มฆวัน โทจันทร"} id={"B6107826"} system={"ระบบงานวิจัยและฐานข้อมูลออนไลน์"}></CardTeam>
          <CardTeam name={"นาย วราเมธ นุ้ยเพียร"} id={"B6108526"} system={"ระบบคืนหนังสือ"}></CardTeam>
          <CardTeam name={"นาย สุภชัย เพ็ชธัมรงค"} id={"B6111427"} system={"ระบบจองเครื่องรับชม Video on Demand"}></CardTeam>
          <CardTeam name={"นาย ณัฐนนท์ รักภักดี"} id={"BB6112936"} system={"ระบบระบบจองห้องค้นคว้า"}></CardTeam>
        </Grid>
      </Content>
    </Page>
  );
};

export default WelcomePage;
