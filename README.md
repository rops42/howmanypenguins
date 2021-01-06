![banner](https://user-images.githubusercontent.com/6611776/103823149-caddeb80-5071-11eb-90b9-b60fb8e44c2d.png)

How many penguins your shitty project kills on `npm install` process.

**How to run it :**
Clone this repository  and perform the `hmp` command in the root folder of your project.
It will detect and read your `package.json` file and provide you de nice score of how many penguins 🐧 it just killed.

After running it you will be allowed to inject your hmp-badge on your readme file to proudly show it to the world.
**Exemples :**
| :knife:🐧 🐧 🐧 | :knife:🐧 🐧 🐧 🐧 🐧  |
|--|--|
| :knife:🐧 🐧 🐧 🐧 🐧 🐧 🐧 🐧  | :knife:🐧 🐧 🐧 🐧 🐧 🐧 🐧 🐧 🐧 🐧 🐧 🐧 🐧  |

    ➜ ./hmp
    Welcome on HowManyPenguin CLI
    I'll install all packages listed in your package.json file and calculate how many penguins you just killed...
    Downloading packages...
    Packages downloaded.
    Counting penguins for 408.1M packages size
    Penguins killed:  11
    🔪 🐧 🐧 🐧 🐧 🐧 🐧 🐧 🐧 🐧 🐧🐧

(I'm actually looking for a designer :heart: to help me build ~10levels badge generator )


**The maths behind that :**
We simply calculate the actual packages size and apply the needed power consumption to download it. It is of course based on your project size too.
Depending on this power usage we can easily and precisely calculate how many cute little penguins it must kill.
(Fully random for now)


If you want to improve this tool please feal free to reach me :heart:
