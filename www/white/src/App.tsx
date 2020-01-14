import React from 'react';
import CssBaseline from '@material-ui/core/CssBaseline';

import './App.css';
import Index from './views/Index';

const App: React.FC = () => {
  return (
    <React.Fragment>
      <CssBaseline />
      <Index />
    </React.Fragment>
  );
}

export default App;
