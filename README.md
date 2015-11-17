# vip

Vim plugin manager

## Philosophy

`vip` **does**:

* Install a plugin from Git repository
* Pin the plugin version to specified version
* Generate installed plugin information with versions
* Update installed plugins
* Run hooks when a plugin updated
* In manner of popular tools such as bundler, carton, or npm

`vip` **does not**:

* Manage `rtp`
* Supports lazy loading
* Takes care of dependencies
