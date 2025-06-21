import { BrowserRouter } from 'react-router-dom';

import './styles/App.css';
import SomeApp from './pages/menu';

function App() {
  return (
    <BrowserRouter>
      <SomeApp/>
    </BrowserRouter>
)
}

export default App;