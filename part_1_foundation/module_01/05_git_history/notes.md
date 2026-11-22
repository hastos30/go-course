- `git log --oneline -- <путь>` — история одного файла

- Три зоны: working tree (стол) → git add → staging area (коробка) → git commit → repository.
- git status: Untracked / Changes not staged / Changes to be committed. Один файл может быть в двух разделах.
- Tracked — Git знает файл. Staged — изменение войдёт в коммит.
- git diff (стол ↔ коробка), git diff --staged (коробка ↔ коммит, проверка перед коммитом), git diff HEAD, git diff A B.
- git restore <файл> — откатить стол к коробке, необратимо. git restore --staged <файл> — вынуть из коробки, безопасно.
- Один коммит — одна мысль, add по пути.
- git log (--oneline, -N, --stat, -- <путь>), git show <коммит>, HEAD~N.
- push → GitHub, fetch → обновить origin/main, pull = fetch + обновить ветку. status не знает о GitHub без fetch.