import {Vue} from 'annotation';
import VueI18n from 'vue-i18n';

import environment from 'environment';
import indonesia from '../locale/indonesia';

Vue.use(VueI18n);

export default new VueI18n({
  locale: environment.LOCALE,
  fallbackLocale: 'id',
  messages: {
    'id': indonesia
  }
});
