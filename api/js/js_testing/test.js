const e = async () => {
    const a = await fetch("https://api.github.com/users/Chris-Coleongco/repos")

    console.log(await a.json())

}

e()
