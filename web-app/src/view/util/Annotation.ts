/**
 * Created by fajar on 3/8/18.
 *
 * Util to provide custom Annotation dependence on
 * 'vue-property-decorator' and vue-class-component'
 * prefer for add custom hook which not provide by vue
 *
 * This file can import by alias 'annotation' or can be
 * import manually (attention the location of this file)
 */

import {Container} from 'typescript-ioc';
import {createDecorator, mixins} from 'vue-class-component';
import {Component, Vue, Provide, Model, Prop, Watch, Emit, Inject} from 'vue-property-decorator';

export * from 'typescript-ioc';
export const IOContainer = Container;

/**
 * Redefine Mixin from vue-class-component as
 * upper case Mixin to prevent ambiguous
 */
export const Mixin = mixins;

/**
 * Redefine Inject from vue-property-decorator as
 * InjectProp to prevent duplicate with typescript-ioc
 */
export const InjectProp = Inject;
export {Component, Vue, Provide, Model, Prop, Watch, Emit}

/** Internal function to merging option into component */
function mergeToComponent(target: any,
                          propertyKey: string,
                          asKey: string = '',
                          invoke: boolean = false,
                          cleanMethod: boolean = true): void {

  return createDecorator(function (component) {
    asKey = asKey === '' ? propertyKey : asKey;
    if (!component[asKey]) {
      component[asKey] = invoke ? target[propertyKey]() : target[propertyKey];

      if (cleanMethod) {
        delete component.methods[propertyKey];
      }
    }
  })(target);
}

/**
 * For html meta tag
 * usage: @MetaTag function() { return object }
 */
export function MetaTag(target: any, propertyKey: string): void {
  mergeToComponent(target, propertyKey, 'metaInfo');
}

/**
 * For component option tasks (vuency)
 * usage: @Task function(t, options) {}
 */
export function Task(target: any, propertyKey: string): void {
  mergeToComponent(target, propertyKey, 'tasks');
}

/**
 * For extend component
 * usage: @Extend function() { return array }
 */
export function Extend(target: any, propertyKey: string): void {
  mergeToComponent(target, propertyKey, 'extends', true, true);
}


/**
 * For extend component
 * usage: @Extend function() { return array }
 */
export function Components(target: any, propertyKey: string): void {
  mergeToComponent(target, propertyKey, 'components', true, true);
}
