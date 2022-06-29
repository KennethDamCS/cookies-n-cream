import React from 'react'
import './App.css';
import DashboardPage from './pages/DashboardPage';
import Footer from './components/Footer';
import Header from './components/header';

function App() {
  return (
    <div className="App">
      <Header />
      <DashboardPage />
      <Footer />
    </div>
  );
}

export default App;
