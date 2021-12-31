interface NavbarProps {
	children: React.ReactNode
	isOpen: boolean
}

function Navbar(props: NavbarProps) {
	const classList = [
		"absolute",
		"lg:static",
		"top-0",
		"left-0",
		"w-72",
		"h-full",
		"bg-slate-400"
	]

	if (!props.isOpen) {
		classList.push("hidden", "lg:block");
	}

	return (
		<div className={classList.join(" ")}>
			{props.children}
		</div>
	)
}

export default Navbar;