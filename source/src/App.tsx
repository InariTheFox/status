import React, { useState } from 'react';
import HomePage from './pages/Home';
import UserPageContainer from './layouts/UserPageContainer';
import { Route, Routes } from 'react-router-dom';
import classNames from 'classnames';

function App() {

  const [isDarkMode, setIsDarkMode] = useState(true);

  const toggleDarkMode = () => {
    setIsDarkMode(!isDarkMode);
  }

  return (
    <div className={classNames("App", { "dark": isDarkMode })}>
      <Routes>
        <Route path="/" element={<UserPageContainer onToggleDarkMode={toggleDarkMode} isDarkMode={isDarkMode}><HomePage /></UserPageContainer>} />
      </Routes>
    </div>
  );
}

export default App;
