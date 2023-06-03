import { Route, Router } from 'preact-router';
import { IntlProvider } from 'preact-i18n';
import { useEffect, useState } from 'preact/hooks';

import './global.css';

import Home from './routes/Home';
import Login from './routes/Login';

import Header from './components/Header';

 
export function App() {
	const [definition, setDefinition] = useState({});

	async function fetchLocale() {
		const response = await fetch(`/i18n/en-US.json`);
		const localeStrings = await response.json();

		setDefinition(localeStrings);
	}

	useEffect(() => {
		

		fetchLocale();
	}, []);

	return (
		<IntlProvider definition={definition}>
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
