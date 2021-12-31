import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faBars, faPlus, faTimes } from '@fortawesome/free-solid-svg-icons';
import Button from "../components/Button";

interface HeaderProps {
	navbarCollapseCallback() : void
	navbarIsOpen : boolean
	classes: JSX.Element[]
}

function Header(props : HeaderProps) {
	return (
		<div className="bg-slate-400 h-20 text-slate-100 text-xl">
			<div className="flex flex-row content-center items-center h-full px-7">
				<Button
					text={<FontAwesomeIcon icon={props.navbarIsOpen ? faTimes : faBars}/>}
					onClick={props.navbarCollapseCallback}
					className="lg:hidden"
				/>
				<div className="mx-auto md:ml-8 lg:ml-0">
					App Name
				</div>
				<Button
					text={<FontAwesomeIcon icon={faPlus}/>}
					onClick={() => {}}
				/>
			</div>
		</div>
	)
}

export default Header;