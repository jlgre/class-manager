interface PageProps {
	children: React.ReactNode
}

function Page(props: PageProps) {
	return(
		<div className="bg-slate-50 text-slate-900 dark:bg-slate-900 dark:text-slate-300 p-5 h-full">
			{props.children}
		</div>
	)
}

export default Page;