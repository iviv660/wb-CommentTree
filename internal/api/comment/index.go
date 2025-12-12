package comment

import (
	"fmt"
	"net/http"
)

func (a *API) index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<!DOCTYPE html>
<html lang="ru">
<head>
  <meta charset="UTF-8">
  <title>Дерево комментариев</title>
  <style>
    body { font-family: sans-serif; margin: 20px; }
    .comment { margin: 4px 0; }
    .children { margin-left: 20px; border-left: 1px dashed #ccc; padding-left: 10px; }
    .meta { font-size: 12px; color: #666; }
    form { margin: 8px 0; }
    input[type="text"] { padding: 4px; width: 300px; }
    textarea { width: 300px; height: 60px; }
    button { padding: 4px 8px; margin-top: 4px; }
  </style>
</head>
<body>
  <h1>Комментарии</h1>

  <section>
    <h2>Поиск</h2>
    <input id="searchQuery" type="text" placeholder="ключевые слова">
    <button id="searchBtn">Искать</button>
  </section>

  <section>
    <h2>Добавить комментарий</h2>
    <form id="commentForm">
      <label>Parent ID (опционально): <input id="parentId" type="text"></label><br>
      <label>Текст:<br><textarea id="commentText"></textarea></label><br>
      <button type="submit">Отправить</button>
    </form>
  </section>

  <section>
    <h2>Дерево комментариев</h2>
    <div id="comments"></div>
  </section>

<script>
async function loadComments() {
  const q = document.getElementById('searchQuery').value;
  const params = new URLSearchParams();
  params.set('limit', '100');
  if (q) params.set('q', q);

  const resp = await fetch('/comments?' + params.toString());
  if (!resp.ok) {
    const errText = await resp.text();
    document.getElementById('comments').innerText =
      'Ошибка загрузки комментариев: ' + resp.status + ' ' + errText;
    return;
  }
  const data = await resp.json();
  const tree = buildTree(data);
  const container = document.getElementById('comments');
  container.innerHTML = '';
  container.appendChild(renderTree(tree, null));
}

function buildTree(comments) {
  const byParent = {};
  for (const c of comments) {
    const pid = c.parent_id === null ? 'root' : String(c.parent_id);
    if (!byParent[pid]) byParent[pid] = [];
    byParent[pid].push(c);
  }
  return byParent;
}

function renderTree(byParent, parentKey) {
  const key = parentKey === null ? 'root' : String(parentKey);
  const items = byParent[key] || [];
  const container = document.createElement('div');

  for (const c of items) {
    const div = document.createElement('div');
    div.className = 'comment';

    const meta = document.createElement('div');
    meta.className = 'meta';
    meta.innerText = 'ID ' + c.id + (c.parent_id !== null ? ' (parent ' + c.parent_id + ')' : '');
    div.appendChild(meta);

    const body = document.createElement('div');
    body.innerText = c.body;
    div.appendChild(body);

    const actions = document.createElement('div');
    const replyBtn = document.createElement('button');
    replyBtn.innerText = 'Ответить';
    replyBtn.onclick = () => {
      document.getElementById('parentId').value = c.id;
      document.getElementById('commentText').focus();
    };
    const deleteBtn = document.createElement('button');
    deleteBtn.innerText = 'Удалить';
    deleteBtn.style.marginLeft = '8px';
    deleteBtn.onclick = () => deleteComment(c.id);
    actions.appendChild(replyBtn);
    actions.appendChild(deleteBtn);
    div.appendChild(actions);

    const children = renderTree(byParent, c.id);
    if (children.childNodes.length > 0) {
      const wrap = document.createElement('div');
      wrap.className = 'children';
      wrap.appendChild(children);
      div.appendChild(wrap);
    }

    container.appendChild(div);
  }

  return container;
}

async function createComment(e) {
  e.preventDefault();
  console.log('submit handler called');

  const parentIdRaw = document.getElementById('parentId').value.trim();
  const text = document.getElementById('commentText').value.trim();
  if (!text) {
    alert('Текст обязателен');
    return false;
  }

  let parent_id = null;
  if (parentIdRaw !== '') {
    const n = Number(parentIdRaw);
    if (Number.isNaN(n)) {
      alert('Parent ID должен быть числом');
      return false;
    }
    parent_id = n;
  }

  const payload = { parent_id, body: text };
  console.log('sending payload', payload);

  const resp = await fetch('/comments', {
    method: 'POST',
    headers: {'Content-Type': 'application/json'},
    body: JSON.stringify(payload),
  });

  if (!resp.ok) {
    const errText = await resp.text();
    alert('Ошибка создания комментария: ' + resp.status + ' ' + errText);
    return false;
  }

  const created = await resp.json();
  console.log('created comment', created);

  document.getElementById('commentText').value = '';
  document.getElementById('parentId').value = '';
  loadComments();
  return false;
}

async function deleteComment(id) {
  if (!confirm('Удалить комментарий #' + id + '?')) return;
  const resp = await fetch('/comments/' + id, { method: 'DELETE' });
  if (!resp.ok) {
    const errText = await resp.text();
    alert('Ошибка удаления: ' + resp.status + ' ' + errText);
    return;
  }
  loadComments();
}

document.getElementById('commentForm').addEventListener('submit', createComment);
document.getElementById('searchBtn').addEventListener('click', (e) => {
  e.preventDefault();
  loadComments();
});

loadComments();
</script>
</body>
</html>
`)
}
