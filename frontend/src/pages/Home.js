import { Helmet } from 'react-helmet';

function Home() {

  return (
    <>
      <Helmet>
        <title>ToDo</title>
      </Helmet>
      <main>
        <div>
            <center><p>This is my ToDo app.</p></center>
        </div>
      </main>
    </>
  );
}

export default Home;