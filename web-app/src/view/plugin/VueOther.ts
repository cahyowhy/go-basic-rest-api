import {Vue} from 'annotation';
import Buefy from 'buefy';
import VImgFallback from 'v-img-fallback';
import 'buefy/dist/buefy.css';

Vue.use(VImgFallback, {
  loading: '/public/images/loading.gif',
  error: '/public/images/no-img-found.png'
});

Vue.use(Buefy);