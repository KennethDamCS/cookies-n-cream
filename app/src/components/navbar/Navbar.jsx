import './Navbar.css'

function Navbar () {

  return (
    <section className="navbar">
      <a href="/dashboard" className="navbar-item">Dashboard</a>
      <a href="/explore" className="navbar-item">Explore</a>
      <a href="/inbox" className="navbar-item">Inbox</a>
      <a href="/settings" className="navbar-item">Settings</a>
      <a href="/login" className="navbar-item">Login</a>
    
  </section>
  )

}

export default Navbar;