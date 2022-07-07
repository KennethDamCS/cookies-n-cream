import Navbar from '../navbar/Navbar';
import Searchbar from '../searchbar/Searchbar';
import React from 'react'
import './Header.css';

function Header () {

  return (
    <section className="ultimate-header">
      <section className="header-top">
        <section className="header-top__logo">
          <a href="/" className="header-logo">PHEROMONE</a>
        </section>
        <section className="header-top__navbar">
          <section className="header-top__navigation">
            <Searchbar />
            <Navbar />
          </section>
        </section>
      </section>
    </section>
  )
}

export default Header;