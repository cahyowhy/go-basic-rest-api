/**
 * Created by fajar on 01/22/2018.
 * Updated by cahyo on 24/10/2018.
 *
 * Vue Default config instance
 * to inject or register anything can be here
 */

import { Vue } from 'annotation';
import Constant from '../config/Constant';

// light component
import CommonImage from '../components/CommonImage.vue';
import EmptyStates from '../components/EmptyStates.vue';
import NavigationBar from '../components/NavigationBar.vue';
import UserBadge from '../components/UserBadge.vue';
import CommonTable from '../components/CommonTable.vue';

Vue.config.productionTip = false;
Vue.config.devtools = false;
Vue.config.silent = true;
Vue.config.errorHandler = (error: any) => console.log(error);

// exporting constant string property to template
Vue.prototype.Constant = Constant;

// register light component
Vue.component('common-image', CommonImage);
Vue.component('empty-states', EmptyStates);
Vue.component('navigation-bar', NavigationBar);
Vue.component('user-badge', UserBadge);
Vue.component('common-table', CommonTable);

// TODO://get rid of this red shit!
// lazy load heavy component
Vue.component('form-login', () => import("../components/FormLogin.vue"));
Vue.component('user-home', () => import("../components/UserHome.vue"));
Vue.component('settings-user', () => import("../components/SettingsUser.vue"));
Vue.component('admin-page', () => import("../components/AdminPage.vue"));
Vue.component('common-editor', () => import("../components/CommonEditor.vue"));

// register all filter in util dir,
// ex: can be use from template using 'param | filterName'
// or method under vue instance using this.$options.filters.filterName
const filters: any = require.context('../util', true, /\.(ts)$/i);
filters.keys().map((key: any) => {
  let filterName: any = key.match(/\w+/)[0];

  if (filterName !== 'Annotation') {
    const filter: any = filters(key);

    Object.keys(filter).map((exportKey) => {
      Vue.filter(exportKey, filter[exportKey])
    });
  }
});
