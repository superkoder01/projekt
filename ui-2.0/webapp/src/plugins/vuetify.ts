// Styles
import '@mdi/font/css/materialdesignicons.css';
import 'vuetify/styles';
import * as components from 'vuetify/lib/components';
import * as directives from 'vuetify/lib/directives';

// Vuetify
import { createVuetify } from 'vuetify';

export default createVuetify(
  // https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
  {
    components,
    directives,
    theme: {
      themes: {
        light: {
          colors: {
            primary: '#b8a041',
            background: '#35495e',
            error: '#d63031',
            info: '#9ae309',
            secondary: '#ffcbff',
            success: '#00cec9',
            surface: '#7fe75c',
            warning: '#2d3436',
            test: '#ffd93f'
          },
          dark: false,
          variables: {}
        }
      }
    }
  }
);
