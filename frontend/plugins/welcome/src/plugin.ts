import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import WatchVideo from './components/WatchVideo'
import SignIn from './components/SignIn'
import Booking from './components/Booking'
import Bookborrow from './components/Bookborrow'
import Bookreturn from './components/Bookreturn'
import Room from './components/Room'
import Research from './components/Research'

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/watch_video', WatchVideo);
    router.registerRoute('/signin', SignIn);
    router.registerRoute('/playlist_video', SignIn);
    router.registerRoute('/Bookborrow', Bookborrow);
    router.registerRoute('/Bookreturn', Bookreturn);
    router.registerRoute('/Research', Research);
    router.registerRoute('/VideoOnDemand', Booking);
    router.registerRoute('/Room', Room);
  },
});
