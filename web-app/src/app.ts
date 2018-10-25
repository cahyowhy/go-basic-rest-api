import { Vue, Inject } from 'annotation';
import CommonService from './view/service/CommonService';
import queryString from 'query-string';
import Turbolinks from 'turbolinks';
import environment from "environment";

import './view/plugin/VueCookie';
import './view/plugin/VueDefault';
import './view/plugin/VueOther';
import i18n from './view/plugin/Vuei18n';

function handleVueDestruction(vue) {
    document.addEventListener('turbolinks:visit', function teardown() {
        vue.$destroy();
        document.removeEventListener('turbolinks:visit', teardown);
    });
}

class App {

    @Inject
    private commonService: CommonService;

    init() {
        // if dev env inject live reload manually
        if (environment['ENV'] !== 'PROD') {
            document.addEventListener('DOMContentLoaded', function () {
                let script = document.createElement('script');
                script.type = 'text/javascript';
                script.src = 'http://localhost:35729/livereload.js';

                document.body.appendChild(script);
            }, false);
        }

        (Date as any).prototype.addDays = function (days) {
            var date = new Date(this.valueOf());
            date.setDate(date.getDate() + days);
            return date;
        };

        const context: App = this;
        let el = '#app';

        Turbolinks.start();
        document.addEventListener('turbolinks:load', () => {
            (window as any).appPreview = new Vue({
                el, i18n,
                beforeMount: function () {
                    if (this.$el.parentNode) {
                        handleVueDestruction(this);
                        this.$originalEl = this.$el.outerHTML;
                    }
                },
                destroyed: function () {
                    this.$el.outerHTML = this.$originalEl;
                },
                computed: {
                    route() {
                        const urlLocation = (window as any).location;
                        const path = urlLocation.pathname;
                        const query = queryString.parse(urlLocation.search) || {};
                        const hash = queryString.parse(urlLocation.hash) || {};

                        return { path, query, hash };
                    }
                }
            });

            context.commonService.app = (window as any).appPreview;
        });
    }
}

const app = new App();
app.init();

export default app;