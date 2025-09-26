#include "webview.h"

#include <stdlib.h>
#include <stdint.h>

#if defined(__linux__)
#include <gtk/gtk.h>
#include <webkit2/webkit2.h>
#endif

struct binding_context {
    webview_t w;
    uintptr_t index;
};

void _webviewDispatchGoCallback(void *);
void _webviewBindingGoCallback(webview_t, char *, char *, uintptr_t);

static void _webview_dispatch_cb(webview_t w, void *arg) {
    _webviewDispatchGoCallback(arg);
}

static void _webview_binding_cb(const char *id, const char *req, void *arg) {
    struct binding_context *ctx = (struct binding_context *) arg;
    _webviewBindingGoCallback(ctx->w, (char *)id, (char *)req, ctx->index);
}

void CgoWebViewDispatch(webview_t w, uintptr_t arg) {
    webview_dispatch(w, _webview_dispatch_cb, (void *)arg);
}

void CgoWebViewSetUserAgent(webview_t w, const char *ua) {
#if defined(__linux__)
    if (w == NULL || ua == NULL || ua[0] == '\0') return;

    GtkWidget *win = (GtkWidget *)webview_get_window(w);
    if (win == NULL) return;

    GtkWidget *child = gtk_bin_get_child(GTK_BIN(win));
    if (child == NULL) return;

    WebKitWebView *wv = WEBKIT_WEB_VIEW(child);

    WebKitSettings *settings = webkit_web_view_get_settings(wv);
    if (settings == NULL) {
        settings = webkit_settings_new();
    }
    webkit_settings_set_user_agent(settings, ua);
    webkit_web_view_set_settings(wv, settings);
#else
    (void)w; (void)ua;
#endif
}

void CgoWebViewBind(webview_t w, const char *name, uintptr_t index) {
    struct binding_context *ctx = calloc(1, sizeof(struct binding_context));
    ctx->w = w;
    ctx->index = index;
    webview_bind(w, name, _webview_binding_cb, (void *)ctx);
}

void CgoWebViewUnbind(webview_t w, const char *name) {
    webview_unbind(w, name);
}
