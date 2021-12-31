import Header from './views/Header';
import Body from './views/Body';
import { useState } from 'react';

function App() {
    const [appNavbarIsOpen, setAppNavbarIsOpen] = useState(false);

    return (
        <div className="h-screen flex flex-col">
            <Header navbarIsOpen={appNavbarIsOpen} navbarCollapseCallback={() => setAppNavbarIsOpen(!appNavbarIsOpen)} classes={[]}/>
            <Body navbarIsOpen={appNavbarIsOpen} groups={[]}/>
        </div>
        );
    }

export default App;
