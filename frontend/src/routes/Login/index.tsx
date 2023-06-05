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
			<div class="background">
				<div class="wrapper">
					<h1>
						<Text id="welcome">Welcome 👋</Text>
					</h1>
					<p>
						<Text id="description">Enter your credentials or create new logisshtickID to continue</Text>
					</p>
					<div class="selector">
						<button id="signIn">
							<Text id="signIn">Sign In</Text>
						</button>
						<button id="createNewAccount">
							<Text id="createNewAccount">Create new account</Text>
						</button>
					</div>

					<form>
						<label id="email">Email</label>
						<input type="text" placeholder="sema_pidoras@proton.me" />
						<label><Text id="password">Password</Text></label>
						<input type="password" placeholder="" />

						<button onClick={() => auth()}><Text id="continue">Continue</Text></button>
					</form>
				</div>
			</div>
        </>
	)
}


export default Login;