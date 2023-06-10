import { Route, Router } from 'preact-router';
import { IntlProvider } from 'preact-i18n';
import { useEffect, useState } from 'preact/hooks';

import './global.css';
import './fonts.css'

import Home from './routes/Home';
import Login from './routes/Login';

import Header from './components/Header';
import Splash from './components/Splash';

 
export function App() {
	const [definition, setDefinition] = useState({});
	//const [splash, setSplash] = useState(true);

	async function fetchLocale() {
		const response = await fetch(`/i18n/pl-PL.json`);
		const localeStrings = await response.json();

		setDefinition(localeStrings);
	}

	useEffect(() => {
		fetchLocale();

		//setTimeout(() => setSplash(false), 1000);

		
	}, []);

	return (
		<IntlProvider definition={definition}>
			<Splash />
			<main >
				<Header />
				<Router>
					<Route path="/" component={Home} />
					<Route path="/login" component={Login} />
				</Router>
			</main>
		</IntlProvider>
	)
}
