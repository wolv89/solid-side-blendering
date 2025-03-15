/* @refresh reload */
import { render } from 'solid-js/web';
import { lazy } from 'solid-js';
import { Router, Route } from '@solidjs/router';

import './index.css';

const App = lazy(() => import("./App"))
const Other = lazy(() => import("./Other"))

const root = document.getElementById('root');

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
  throw new Error(
    'Root element not found. Did you forget to add it to your index.html? Or maybe the id attribute got misspelled?',
  );
}

const routes = () => (
  <Router>
    <Route path="/other" component={Other} />
    <Route path="/" component={App} />
  </Router>
)

render(routes, root);
