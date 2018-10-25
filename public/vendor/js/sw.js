/*
 *  Kita akan ubah nama
 *  pada cacheVersion
 *  supaya browser tahu bahwa
 *  ada perubahan terbaru
 */
const cacheVersion = 'todo-app-2';
const filesToCache = [
  '/',
  '/public/css/app.css',
  '/public/js/app.js',
  '/public/vendor/js/jquery.min.js',
  '/public/vendor/css/non-priority.vendor.css',
  '/public/vendor/js/non-priority.vendor.js',
  '/public/vendor/js/jquery.migrate.min.js',
];

self.addEventListener('install', function(event) {
  console.log('start fetch');
  event.waitUntil(
    caches.open(cacheVersion)
      .then(function(cache) {
        return cache.addAll(filesToCache)
      })
  )
});

self.addEventListener('fetch', function(event) {
  event.respondWith(
    caches.match(event.request)
      .then(function(res) {
        if (res) return res;

        return fetch(event.request);
      })
  );
});

self.addEventListener('activate', function(event) {
  event.waitUntil(
    caches.keys().then(function(cacheNames) {
      return Promise.all(
        cacheNames
          .filter(function(cacheName) {
            return cacheName !== cacheVersion;
          })
          .map(function(cacheName) {
            caches.delete(cacheName);
          })
      );
    })
  );
});

self.addEventListener('message', function(event) {
  if (event.data.action === 'skipWaiting') {
    self.skipWaiting();
  }
});
