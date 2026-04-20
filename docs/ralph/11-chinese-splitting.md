# Chinese Literacy Speedrun II: Character Cyclotron
**Signal**: Handling multi-byte Unicode (e.g., Chinese characters)
**Domain experiment**: Test `Split` with UTF-8 byte splitting for Chinese characters
**Next failing test**: Split("中文", "") = ["中", "文"], want ["中", "文"] (current test passes, but needs explicit validation)
