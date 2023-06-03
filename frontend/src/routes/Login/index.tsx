import './styles.css'; 
import { Text } from 'preact-i18n'; 

function Login() {
	async function auth () {
		
		const response = await fetch(`${import.meta.env.VITE_API_URL}/login`, {

		});
		
		console.log(response);


	}

	return (
		<>
			<div class="wrapper">
				<h1>
					<Text id="authorize">Authorize</Text>
				</h1>
				<form>
					<label>Email</label>
					<input type="text" placeholder="sema_pidoras@proton.me" />
					<label><Text id="password">Password</Text></label>
					<input type="password" placeholder="" />

					<button onClick={() => auth()}><Text id="signIn">Sign In</Text></button>
				</form>
			</div>
        </>
	)
}


export default Login;