import { Link } from 'preact-router/match';

import './styles.css';

function Header() {
	return (
		<header style="display: none;">
			<nav>	
				<Link activeClassName="selected" href="/" aria-label="Home">Home</Link>
				<Link activeClassName="selected" href="/settings" aria-label="Settings">Settings</Link>
			</nav>
		</header>
	);
}

export default Header;