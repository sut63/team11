import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import WatchVideo from './components/WatchVideo'
import SignIn from './components/SignIn'
import Booking from './components/Booking'
<<<<<<< HEAD
import ResearchCreate from './components/ResearchCreate'
=======
import Bookborrow from './components/Bookborrow'
>>>>>>> 498be5f28d5dab1a99204d7659387338e06bc606


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/watch_video', WatchVideo);
    router.registerRoute('/signin', SignIn);
    router.registerRoute('/playlist_video', SignIn);
<<<<<<< HEAD
    router.registerRoute('/researchcreate', ResearchCreate);
=======
    router.registerRoute('/Bookborrow', Bookborrow);
>>>>>>> 498be5f28d5dab1a99204d7659387338e06bc606
  },
});
