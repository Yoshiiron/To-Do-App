import { useState } from 'react';
import { Routes, Route, NavLink, useLocation } from 'react-router-dom';

import Home from './Home';      
import ItemList from './Issues';
import CreateForm from './CreateIssue';
import Modal from './modal';

import '../styles/App.css'


function SomeApp() {
  const [menuOpen, setMenuOpen] = useState(false);
  const [modalOpen, setModalOpen] = useState(false);
  const location = useLocation();

  const toggleMenu = () => setMenuOpen(prev => !prev);
  const closeMenu = () => setMenuOpen(false);

  return (
    <>
      <header className='menu'>
        {!menuOpen && (
          <button className="menu-button" onClick={toggleMenu}>
            ☰ Menu
          </button>
        )}
        {!menuOpen && location.pathname === '/issues' && (<button 
          className='menu-button'
          onClick={() => setModalOpen(true)}
        >
          Create an Issue
        </button>)}
      </header>

      <nav className={`side-menu ${menuOpen ? 'open' : ''}`}>
        <NavLink to="/" end className={({ isActive }) => (isActive ? 'active' : '')} onClick={closeMenu}>
          Main
        </NavLink>
        <NavLink to="/issues" className={({ isActive }) => (isActive ? 'active' : '')} onClick={closeMenu}>
          Issues
        </NavLink>
      </nav>

      <main className={menuOpen ? 'blurred' : ''}>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/issues" element={<ItemList />} />
        </Routes>
      </main>

      <Modal isOpen={modalOpen} onClose={() => setModalOpen(false)}>
        <center><h2 className='create'>Create an Issue</h2></center>
        {<CreateForm/>}
      </Modal>
      </>
  );
}

export default SomeApp;