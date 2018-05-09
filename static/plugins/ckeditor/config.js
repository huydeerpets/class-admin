/**
 * @license Copyright (c) 2003-2018, CKSource - Frederico Knabben. All rights reserved.
 * For licensing, see https://ckeditor.com/legal/ckeditor-oss-license
 */

CKEDITOR.editorConfig = function( config ) {
	config.height=600;
	config.removePlugins='language,save,newpage,print,smiley,anchor,pagebreak,iframe,about';
	config.extraPlugins='html5video';
    config.filebrowserImageUploadUrl = "/spider/upload?type=image";
    config.filebrowserFlashUploadUrl = '/spider/upload';
};
