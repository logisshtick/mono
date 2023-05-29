import { Route, Router } from 'preact-router';
import { IntlProvider } from 'preact-i18n';
import { useEffect, useState } from 'preact/hooks';

import Home from './routes/Home';


export function App() {
	const [definition, setDefinition] = useState({});

	async function fetchLocale() {
		let definition = await fetch(`${import.meta.env.VITE_API_URL}/translations`);
		setDefinition(definition);
	}

	useEffect(() => {
		fetchLocale();
	}, []);

	return (
		<IntlProvider scope="weather" definition={definition}>

		<Router>
			<Route path="/" component={Home} />
			
		</Router>

		</IntlProvider>
	)
}
