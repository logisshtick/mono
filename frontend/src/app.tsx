import { Route, Router } from 'preact-router';
import { IntlProvider } from 'preact-i18n';
import { useEffect, useState } from 'preact/hooks';

import Home from './routes/Home';
import Settings from './routes/Settings';
import Login from './routes/Login';

import Header from './components/Header';


export function App() {
	const [definition, setDefinition] = useState({});

	async function fetchLocale() {
		const definition = await fetch(`${import.meta.env.VITE_API_URL}/translations`);
		setDefinition(definition);
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
					<Route path="/settings" component={Settings} />
					<Route path="/login" component={Login} />
				</Router>
			</main>
		</IntlProvider>
	)
}
