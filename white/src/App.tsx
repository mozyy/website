import React from 'react';
import CssBaseline from '@material-ui/core/CssBaseline';
// import 'normalize.css';
import './App.css';
import White from './components/White';

const App: React.FC = () => {
  return (
    <React.Fragment>
      <CssBaseline />
      <div className="App">
        <White />
      </div>
    </React.Fragment>
  );
}

export default App;
