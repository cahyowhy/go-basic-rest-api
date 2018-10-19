/**
 * Created by fajar on 01/22/2018.
 *
 * Vue Default config instance
 * to inject or register anything can be here
 */

import {Vue} from 'annotation';
import Constant from '../config/Constant';

Vue.config.productionTip = false;
Vue.config.devtools = false;
Vue.config.silent = true;
Vue.config.errorHandler = (error: any) => console.log(error);

// exporting constant string property to template
Vue.prototype.Constant = Constant;

// register all component, ex: NagivationBar => accessible to navigation-bar
const components: any = require.context('../components', true, /\.(vue)$/i);
components.keys().map((key: any) => {
  let componentName: any = key.match(/\w+/)[0];

  // check has sub folder
  if ((key.match(/\//g) || []).length > 1) {
    componentName = key.split('/');
    componentName = componentName[componentName.length - 1];
    componentName = componentName.match(/\w+/)[0];
  }

  Vue.component(componentName, components(key));
});

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
