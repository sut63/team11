import { createPlugin } from '@backstage/core';
import SignIn from './components/SignIn'
import Booking from './components/Booking'
import Bookborrow from './components/Bookborrow'
import Bookreturn from './components/Bookreturn'
import Room from './components/Room'
import Research from './components/Research'
import Book from './components/Book'
import Library from './components/Library'
import SearchBooking from './components/SearchBooking'
import SearchBookreturn from './components/SearchBookreturn'
import SearchResearch from './components/SearchResearch'
import SearchBook from './components/SearchBook'
import SearchBookborrow from './components/SearchBookborrow'

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Library);
    router.registerRoute('/SignIn', SignIn);
    router.registerRoute('/Bookborrow', Bookborrow);
    router.registerRoute('/Bookreturn', Bookreturn);
    router.registerRoute('/Research', Research);
    router.registerRoute('/VideoOnDemand', Booking);
    router.registerRoute('/Room', Room);
    router.registerRoute('/Book', Book);
    router.registerRoute('/SearchBooking', SearchBooking);
    router.registerRoute('/SearchBookreturn', SearchBookreturn);
    router.registerRoute('/SearchResearch', SearchResearch);
    router.registerRoute('/SearchBook', SearchBook);
    router.registerRoute('/SearchBookborrow', SearchBookborrow);
  },
});
