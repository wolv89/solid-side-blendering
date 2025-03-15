import { createResource, Show } from 'solid-js';
import styles from './App.module.css';

const fetchEndpoint = async () => {
  const response = await fetch(`http://blewb.build:8338/endpoint`)
  return response.json()
}

function App() {

  const [ep] = createResource(fetchEndpoint)

  return (
    <div class={styles.App}>
      <p>Hello world</p>
      <Show when={ep.loading}>
        <p>Loading...</p>
      </Show>
      <Show when={ep()}>
        <p>{JSON.stringify(ep())}</p>
      </Show>
    </div>
  );
}

export default App;
