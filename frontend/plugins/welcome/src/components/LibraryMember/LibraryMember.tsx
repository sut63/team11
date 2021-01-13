import React from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Typography,
  Grid,
  List,
  ListItem,
  ListItemText,
  Link,
  makeStyles,
} from '@material-ui/core';
import {
  Content,
  InfoCard,
  Header,
  HomepageTimer,
  Page,
  pageTheme,
  ContentHeader,
  Button,
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
    margin: theme.spacing(1, 1, 0, 0),
  },
  img: {
    margin: 'auto',
    display: 'block',
    maxWidth: '100%',
    maxHeight: '100%',
  },
}));
const WelcomePage = () => {
  const classes = useStyles();
  return (
    <Page theme={pageTheme.home}>
      <Header
        title={'Welcome to Library System'}
        subtitle="The Center for Library Resources and Educational Media"
      >
        <HomepageTimer />
      </Header>
      <Content>
        <ContentHeader title="Getting Started">
          <div>
          <Button
            type="submit"
            fullWidth
            className={classes.submit}
            variant="contained"
            color="primary"
            to="/"
          >
            Sign In
          </Button>
          </div>
        </ContentHeader>

        <Grid container>
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
          <Grid item xs={12} md={6}>

            <InfoCard title="What Now?">
              <Typography variant="body1" gutterBottom>
                You now have a running instance of Backstage!&nbsp;
                <span role="img" aria-label="confetti">
                  🎉
                </span>
                &nbsp;Let's make sure you get the most out of this platform by
                walking you through the basics.
              </Typography>
              <Typography variant="h6" gutterBottom>
                The Setup
              </Typography>
              <Typography variant="body1" paragraph>
                Backstage is put together from three base concepts: the core,
                the app and the plugins.
              </Typography>
              <List>
                <ListItem>
                  <ListItemText primary="The core is responsible for base functionality." />
                </ListItem>
                <ListItem>
                  <ListItemText primary="The app provides the base UI and connects the plugins." />
                </ListItem>
                <ListItem>
                  <ListItemText
                    primary="The plugins make Backstage useful for the end users with
                  specific views and functionality."
                  />
                </ListItem>
              </List>
              <Typography variant="h6" gutterBottom>
                Build Your Plugins
              </Typography>
              <Typography variant="body1" paragraph>
                We suggest you either check out the documentation for{' '}
                <Link href="https://github.com/spotify/backstage/blob/master/docs/getting-started/create-a-plugin.md">
                  creating a plugin
                </Link>{' '}
                or have a look in the code for the{' '}
                <Link component={RouterLink} to="/explore">
                  existing plugins
                </Link>{' '}
                in the directory{' '}
                <Link href="https://github.com/spotify/backstage/tree/master/plugins">
                  <code>plugins/</code>
                </Link>
                .
              </Typography>
            </InfoCard>
          </Grid>
          <Grid item>
            <InfoCard title="Quick Links">
              <List>
                <ListItem>
                  <Link href="https://backstage.io">backstage.io</Link>
                </ListItem>
                <ListItem>
                  <Link href="https://github.com/spotify/backstage/blob/master/docs/getting-started/create-a-plugin.md">
                    Create a plugin
                  </Link>
                </ListItem>
                <ListItem>
                  <Link href="/explore">Plugin gallery</Link>
                </ListItem>
              </List>
            </InfoCard>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
};

export default WelcomePage;
